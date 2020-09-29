package main

import (
	"fmt"
)

func main() {
	result := validPalindrome("abca")
	fmt.Println(result)
	result = validPalindrome("aba")
	fmt.Println(result)
	result = validPalindrome("acbca")
	fmt.Println(result)
	result = validPalindrome("abcdddcba")
	fmt.Println(result)
	result = validPalindrome("ejhrpskcrbgvienptzackhfdymdmqmrwhvautlbpxgzyehzzgvqofcxismccbptkqompkrfxytawjitmygwsotcrcuprvaqmcaqixuggsjtjrpveecunmfhcucvoawduyjjvckfjnubdagikmloqsvgvuwnrzwauntdwttsyhwzhrjvpoyelwyyltlocpxnrxdtsebsqfrsryfderhdtahaadrzobmmgdisnienvrckuiofetrcqhcffsknmhtajkfzcucxeunrbypbtfqyqqrxowbarvreyetcmwirrsqfdyfyxchidumsnbvzcceqyytmrrmdglbifktnrpdqrenhcweaeluoqucsfhazdqxsfjiksnzcrclttweabeezddwztdgvfgxndvsebjhxigbvhhismpgqetroaevcyjwmrkjbrctwgqeydjzjupjdlyeneknkkqsjvumeybbakawsnfajkrtjemmkvamtzxitzwwrobvmucndtbmsupkflovbyudepiqocyjcoetawmiumfngqaioylsrlokxwtoifeldqneopwkduudxzqksbvtkajykvxrioxxcpatbjndorhnihbmnwkdhcuseoobesgqdrmiaiasmuexwkihbxmdxirzntxrygfoxwpznhjyhnexidndudcncyvtbwciefalioezdilddnezahomikezyswnjarnbmbspbriaqyxeafbnigwvbwrpfsuibfsmvonmcowgmcgsyrghaeqobofnunjvlzicnhiprpbbkajvrwkyoanpkpfhhmavxfmewytthxftgupapezhbljowfrjroynvngnqkcrgtyvdccimuhqtuqyxpsabminpptspzzntalolrvhwpxzuoqruriylptxkjokhbbwsyhjdnbivxawjsguizxhjiiptkxkqyamvabwjiqbzfugycudjvacufqauguisffupefhzrmxrmvwzwvzpbigbrbboobbrbgibpzvwzwvmrxmrzhfepuffsiuguaqfucavjducygufzbqijwbavmayqkxktpiijhxziugsjwaxvibndjhyswbbhkojkxtplyirurqouzxpwhvrlolatnzzpstppnimbaspxyqutqhumiccdvytgrckqngnvnyorjrfwojlbhzepapugtfxhttywemfxvamhhfpkpnaoykwrvjakbbprpihncizlvjnunfoboqeahgrysgcmgwocmnovmsfbiusfprwbvwginbfaexyqairbpsbmbnrajnwsyzekimohazenddlidzeoilafeicwbtvycncdudndixenhyjhnzpwxofgyrxtnzrixdmxbhikwxeumsaiaimrdqgsebooesuchdkwnmbhinhrondnjbtapcxxoirxvkyjaktvbskqzxduudkwpoenqdlefiotwxkolrslyoiaqgnfmuimwateocjycoqipeduyfbvolfkpusmbtdncumvborwwztixztmavkmmejtrkjafnswakabbyemuvjsqkknkeneyldjpujzjdyeqgwtcrbjkrmwjycveaorteqgpmsihhvbgixhjbesvdnxgfvgdtzwddzeebaewttlcrcznskijfsxqdzahfscuqouleaewchnerqdprntkfiblgdmrrmtyyqecczvbnsmudihcxyfydfqsrriwmcteyervrabwoxrqqyqftbpybrnuexcuczfkjathmnksffchqcrtefoiukcrvneinsidgmmbozrdaahatdhredfyrsrfqsbestdxrnxpcoltlyywleyopvjrhzwhysttwdtnuawzrnwuvgvsqolmkigadbunjfkcvjjyudwaovcuchfmnuceevprjtjsgguxiqacmqavrpucrctoswgymtijwatyxfrkpmoqktpbccmsixcfoqvgzzheyzgxpbltuavhwrmqmdmydfhkcaztpneivgbrcksprhje")
	fmt.Println(result)
}

func validPalindrome(s string) bool {

	//sl := len(s)
	//mid := sl / 2
	//if sl%2 == 1 {
	//	mid += 1
	//}

	reverseS := ""
	for _, v := range s {
		reverseS = string(v) + reverseS
	}
	if s == reverseS {
		return true
	}

	for index, _ := range s {
		s2 := s[:(index)] + s[(index+1):]
		reverseS := ""
		for _, v := range s2 {
			reverseS = string(v) + reverseS
		}
		if s2 == reverseS {
			return true
		}
	}

	return false
}
