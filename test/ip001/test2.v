func readFile(readChannel chan string) {
    f, err := os.Open("1.txt")
    if err != nil {
        panic(fmt.Sprintf("open file error:%s", err.Error()))
    }
    //移动到文件末尾
    f.Seek(0, os.SEEK_END)
    reader := bufio.NewReader(f)
    for {
        line, err := reader.ReadBytes('\n')
        fmt.Println(err)
        if err == io.EOF {
            time.Sleep(time.Second)
            continue
        } else if err != nil {
            panic("ReadBytes error:" + err.Error())
        }

        lineStr := strings.TrimSpace(string(line))
        readChannel <- lineStr
    }
}