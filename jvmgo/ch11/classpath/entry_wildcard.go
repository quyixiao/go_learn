package classpath

import "os"
import "path/filepath"
import "strings"

//然后调用filepath包 的Walk()函数遍历baseDir创建ZipEntry。Walk()函数的第二个参数 也是一个函数，
//了解函数式编程的读者应该一眼就可以认出这种用 法(即函数可作为参数)。
func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1] // remove *
	compositeEntry := []Entry{}

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}

	//遍历一个目录下的所有文件
	filepath.Walk(baseDir, walkFn)
	return compositeEntry
}
