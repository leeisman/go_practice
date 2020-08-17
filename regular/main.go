package main

import (
	"log"
	"regexp"
)

func main() {
	url := "https://www.qchannel03.cn/1.gif?domain=k.sina.cn&url=-&title=%E6%A5%BC%E4%B8%8A%E9%82%BB%E5%B1%85%E9%A3%98%E7%AA%97%E6%94%B9%E9%98%B3%E5%8F%B0%EF%BC%8C%E7%AB%9F%E7%84%B6%E7%A7%81%E8%87%AA%E5%B0%86%E6%89%BF%E9%87%8D%E6%9F%B1%E7%84%8A%E6%8E%A5%E5%9C%A8%E6%A5%BC%E4%B8%8B%E9%82%BB%E5%B1%85%E5%AE%B6%EF%BC%81&referrer=-&sh=1080&sw=1920&cd=24&lang=en-US&account=SinaNews&channel=compony&point=H5&platform=pc&undefined=undefined&jmid=-&ts=1596174556727"
	match, _ := regexp.MatchString(`\.webm|\.mkv|\.x-matroska|\.flv|\.vob|\.ogv|\.avi|\.mp4|\.mpg|\.mpeg|\.mov|\.quicktime|\.wmv|\.rm|\.vnd.rn-realmedia|\.rmvb|\.vnd.rn-realmedia-vbr|\.m4v|\.3gp|\.m3u8`, url)
	log.Print("match: ",match)
}
