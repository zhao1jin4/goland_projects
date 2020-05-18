package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)
func  main()  {
	var myPath string ="D:/tmp/my.txt"
	var myDir string ="d:/tmp/aa"

	//f,err:=os.Open(myPath) //默认是只读的
	f,err:=os.OpenFile(myPath,os.O_APPEND,os.ModePerm)//os.O_APPEND,os.O_WRONLY | os.O_RDONLY
	if err!=nil {
		fmt.Println(err)
		if ins,ok:=err.(*os.PathError);ok {
			fmt.Printf("打开文件%s 错误原因为%s,OP=%s",ins.Path,ins.Err,ins.Op)
		}
		return;
	}
	fmt.Println(f.Name())
	buf:=make([]byte,64,64)
	f.Seek(6,io.SeekStart)//路过前n字节个开始读（如单数有可能中文切了)，中文不行？？？
	//io.SeekCurrent , io.SeekEnd
	for {
		len,err:=f.Read(buf)
		//f.ReadAt()//从指定位置来读
		if len==0  || err==io.EOF {
			break;
		}
		fmt.Print(string(buf[:len]))
	}

	defer f.Close()

	fileInfo,err:=os.Stat(myPath)
	if err!=nil {
		fmt.Println(err)
		return;
	}
	fmt.Printf("IdDir=%t,size=%d,modTime=%s\n",fileInfo.IsDir(),fileInfo.Size(),fileInfo.ModTime())
	fmt.Printf("mod=%s\n",fileInfo.Mode())
	fmt.Printf("IsAbs=%t\n",filepath.IsAbs(myPath)) //是否绝对路径,filepath.Abs()得到绝对路径

	fmt.Printf("Join=%s\n",path.Join(myPath,".."))

	os.Mkdir(myDir,os.ModePerm) //MkdirAll

	f1,err:=os.Create(myDir+"/myfile.txt")
	if err!=nil {
		fmt.Println("建立文件错误，做删除",err)
		os.Remove(myDir+"myfile.txt")//os.RemoveAll()
		return;
	}
	fmt.Println("向文件写内容")
	f1.WriteString("abcde")
	f1.Write([]byte("ABCD"))
	//f1.WriteAt()//指定位置写
	defer f1.Close()

	//copyFile("D:/tmp/my.txt","D:/tmp/my2.txt");//自已的方法
	//copyFileInMem("D:/tmp/my.txt","D:/tmp/my2.txt");//自已的方法
	//copyUseBufio("D:/tmp/my.txt","D:/tmp/my2.txt");
	recursiveShowDir("D:/tmp/",1)


	fmt.Print("请输入")
	in:=bufio.NewReader(os.Stdin) //支持多个空格分隔都能读到
	str,_:=in.ReadString('\n');
	fmt.Print(str)


}
func copyFile(fromFile string,toFile string)(int64,error){
	from,err:=os.OpenFile(fromFile,os.O_RDONLY,os.ModePerm)
	if err!=nil { //每个文件API调用都要if判断错误，确实不太好
		return 0,err;
	}

	to,err:=os.OpenFile(toFile,os.O_WRONLY|os.O_CREATE,os.ModePerm)
	if err!=nil {
		return 0,err;
	}
	buf:=make([]byte,64,64)
	return io.CopyBuffer(to,from,buf)
	//return io.Copy(to,from)
}
func copyFileInMem(fromFile string,toFile string)(int,error) {
	r1:=strings.NewReader("ABC123中文")
	data,err:=ioutil.ReadAll(r1)
	fmt.Printf("%s\n",data)
	//--
	tmpfile,err:=ioutil.TempFile("d:/tmp/","checkbill_*.txt")//会把*替换为随机数
	defer os.Remove(tmpfile.Name())
	defer  tmpfile.Close()
	tmpfile.Write(data);

	//===
	bs,err:=ioutil.ReadFile("D:/tmp/my.txt") //一次性读入内存，不适合文件过大,源码是调用的readAll方法
	if err!=nil {
		return 0,err;
	}
	os.Create(toFile)
	err=ioutil.WriteFile("D:/tmp/my2.txt",bs,os.ModePerm)
	if err!=nil {
		return 0,err;
	}
	return len(bs),nil
}
func copyUseBufio(fromFile string,toFile string)(int,error){
	from,err:=os.OpenFile(fromFile,os.O_RDONLY,os.ModePerm)
	if err!=nil {
		return 0,err;
	}
	defer  from.Close()
	reader:=bufio.NewReader(from)
	//---
	//buf:=make([]byte,64,64)
	//for {
	//	len,err:=reader.Read(buf)
	//	if len==0  || err==io.EOF {
	//		break;
	//	}
	//	fmt.Print(string(buf[:len]))
	//}
	//---按行读
	for{
		//data,flag,err:=reader.ReadLine()//底层的
		//fmt.Printf("flag=%t,err=%s,data=%s",flag,err,string(data))
		//---
		data,err:=reader.ReadString('\n') //还有 reader.ReadBytes('\n') ，reader.ReadByte()
		if err == io.EOF {
			break;
		}
		fmt.Printf("err=%s,data=%s",err,string(data))
	}
	//--写
	to,err:=os.OpenFile(toFile,os.O_WRONLY|os.O_CREATE,os.ModePerm)
	if err!=nil {
		return 0,err;
	}
	defer to.Close()

	writer:=bufio.NewWriter(to)
	len,err:=writer.WriteString("hello中文")
	writer.Flush()//必须手工调用写缓冲
	return len,err
}
func recursiveShowDir(dir string,level int)(int,error){
	fileInfos,err:=ioutil.ReadDir(dir)//相当于在目录下 ls
	if(err!=nil){
		return 0,err;
	}
	tree:="|-"
	for i:=0;i<level;i++ {
		tree="| "+tree
	}
	for _,item:= range fileInfos {
		fmt.Printf("%s %s/%s\n",tree,dir,item.Name())
		if(item.IsDir()){
			recursiveShowDir(dir+"/"+item.Name(),level+1)
		}
	}
	return 0,nil
}