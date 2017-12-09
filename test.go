package main;

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"os/exec"
)
import b64 "encoding/base64";

// askForConfirmation uses Scanln to parse user input. A user must type in "yes" or "no" and
// then press enter. It has fuzzy matching, so "y", "Y", "yes", "YES", and "Yes" all count as
// confirmations. If the input is not recognized, it will ask again. The function does not return
// until it gets a valid response from the user. Typically, you should use fmt to print out a question
// before calling askForConfirmation. E.g. fmt.Println("WARNING: Are you sure? (yes/no)")
func askForConfirmation() bool {
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}
	okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	nokayResponses := []string{"n", "N", "no", "No", "NO"}
	if containsString(okayResponses, response) {
		return true
	} else if containsString(nokayResponses, response) {
		return false
	} else {
		fmt.Printf("[Yes/No]:")
		return askForConfirmation()
	}
}
func askForWhich() bool {
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}
	pocketmine := []string{"1"}
	able := []string{"2"}
	if containsString(pocketmine, response) {
		return true
	} else if containsString(able, response) {
		return false
	} else {
		fmt.Println("[Yes/No]:")
		return askForConfirmation()
	}
}
// You might want to put the following two functions in a separate utility package.

// posString returns the first index of element in slice.
// If slice does not contain element, returns -1.
func posString(slice []string, element string) int {
	for index, elem := range slice {
		if elem == element {
			return index
		}
	}
	return -1
}

// containsString returns true iff slice contains element
func containsString(slice []string, element string) bool {
	return !(posString(slice, element) == -1)
}
func main() {
	fmt.Print("\033[H\033[2J");
	fmt.Println(getMS());
	fmt.Println("이것을 이용하는 사람들중 엠코 관리자가 있다면 부탁하나 해도 될까요?");
	fmt.Println("저같은 경우는 제 주민등록번호를 모르고 부모님께 물어보기도 곤란한데 본인인증때문에 엠코 가입을 못하고 있습니다;;");
	fmt.Println("엠코 본인인증 좀 풀어주세요!!");
	if !cmdExists("git") {
		fmt.Println("git이 설치되지 않았습니다.");
		fmt.Printf("설치하시겠습니까?[Yes/No]:");
		if askForConfirmation() {
			if cmdExists("apt-get") {
				fmt.Print("\033[H\033[2J");
				fmt.Println(getMS());
				fmt.Println("진행상황은 확인할수 없으나 창을 닫지 마시고 아래 절차에 따라 주세요");
				cmd := "sudo"
				args := []string{"apt-get", "-y", "install", "git"}
				if err := exec.Command(cmd, args...).Run(); err != nil {
					fmt.Fprintln(os.Stderr, err)
					os.Exit(1)
				}
				fmt.Print("\033[H\033[2J");
				fmt.Println(getMS());
			} else {
				fmt.Print("\033[H\033[2J");
				log.Fatal("apt 패키지 관리자를 이용하여 설치를 시도하였으나 실패하였습니다 다른 방법으로 git을 설치하여 주십시오");
				os.Exit(1);
			}
		}
	}
	if is64(){
		if _, err := os.Stat("pm"); !os.IsNotExist(err) {
			fmt.Printf("설치하려는 폴더 [pm] 에 폴더가 이미 있습니다, 덮어쓰시겠습니까?[Yes/No]:");
			var _bool = askForConfirmation();
			if _bool {
				os.RemoveAll("pm");
			} else {
				
				os.Exit(0);
			}
		}
		os.MkdirAll("pm", os.ModePerm);
		fmt.Println("1.ABLE-MP(권장)");
		fmt.Println("2.Pocketmine-MP");
		fmt.Printf("선택[1/2]:");
		if askForWhich() {
			fmt.Print("\033[H\033[2J");
			fmt.Println(getMS());
			installABLE("pm");
		} else {
			fmt.Print("\033[H\033[2J");
			fmt.Println(getMS());
			installPM("pm");
		}
		fmt.Println("PHP다운로드 중....");
		os.MkdirAll("pm/bin", os.ModePerm);
		cmd := "git";
		args := []string{"clone", "https://github.com/AbleUnion/pocketmine-php-linux.git", "pm/bin"};
		if err := exec.Command(cmd, args...).Run(); err != nil {
			fmt.Fprintln(os.Stderr, err);
			os.Exit(1);
		}
		fmt.Println("퍼미션 변경 중");
		if err := exec.Command("chmod", "777", "-R","pm").Run(); err != nil {
			fmt.Fprintln(os.Stderr, err);
			os.Exit(1);
		}
		fmt.Println("설치완료!");
		os.Exit(0);
	} else {
		log.Fatal("64비트 운영체제가 아닙니다 64비트 운영체제를 설치하시거나 호스팅 관리자에게 문의해주세요(pm구동기는 64비트를 지원하지 않습니다)");
	}
}
func is64() bool {
	const PtrSize = 32 << uintptr(^uintptr(0)>>63)
	if strconv.IntSize == 64 {
		return true;
	} else {
		return false;
	}
}
func getMS() string{
	sDec, _  := b64.StdEncoding.DecodeString("ICAgICAgIF98ICBffCAgICAgICAgX3wgIF98ICAgICAgICAgICAgICAgICAgICAgICAgICAgIF98ICAgICAgICAgICAgDQogICBffF98X3wgICAgICAgIF98X3xffCAgICAgICAgX3xffF98ICAgIF98X3wgICAgICBffF98X3wgICAgX3xffCAgICANCiBffCAgICBffCAgX3wgIF98ICAgIF98ICBffCAgX3xffCAgICAgIF98ICAgIF98ICBffCAgICBffCAgX3wgICAgX3wgIA0KIF98ICAgIF98ICBffCAgX3wgICAgX3wgIF98ICAgICAgX3xffCAgX3wgICAgX3wgIF98ICAgIF98ICBffCAgICBffCAgDQogICBffF98X3wgIF98ICAgIF98X3xffCAgX3wgIF98X3xffCAgICAgIF98X3wgICAgICBffF98X3wgICAgX3xffCAgICANCiAgICAgICAgICAgX3wgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIA0KICAgICAgICAgX3w=");
	return string(sDec);
}
func cmdExists(cmd string) bool {
	if _, err := os.Stat("/bin/" + cmd); !os.IsNotExist(err) {
		return true;
	}
	if _, err := os.Stat("/sbin/"+ cmd); !os.IsNotExist(err) {
		return true;
	}
	if _, err := os.Stat("/usr/bin/"+ cmd); !os.IsNotExist(err) {
		return true;
	}
	if _, err := os.Stat("/usr/sbin/"+ cmd); !os.IsNotExist(err) {
		return true;
	}
	return false;
}
func installPM(dir string) {
	cmd := "git";
	fmt.Println("pocketmine 다운로드중...");
	args := []string{"clone", "https://github.com/PocketMine/PocketMine-MP.git", dir};
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err);
		os.Exit(1);
	}
	fmt.Println("pocketmine다운로드 완료,raklib다운로드중...");
	args0 := []string{"clone", "https://github.com/PocketMine/RakLib.git", dir + "/src/raklib"};
	if err := exec.Command(cmd, args0...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err);
		os.Exit(1);
	}
	fmt.Println("pocketmine,raklib다운로드 완료,spl다운로드중...");
	args1 := []string{"clone", "https://github.com/PocketMine/PocketMine-SPL.git", dir + "/src/spl"};
	if err := exec.Command(cmd, args1...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err);
		os.Exit(1);
	}
}
func installABLE(dir string) {
	cmd := "git";
	fmt.Println("able-mp 다운로드중...");
	args := []string{"clone", "https://github.com/PocketMine/PocketMine-MP.git", dir};
	if err0 := exec.Command(cmd, args...).Run(); err0 != nil {
		fmt.Fprintln(os.Stderr, err0);
		os.Exit(1);
	}
	fmt.Println("able-mp다운로드 완료,raklib다운로드중...");
	args0 := []string{"clone", "https://github.com/PocketMine/RakLib.git", dir + "/src/raklib"};
	if err1 := exec.Command(cmd, args0...).Run(); err1 != nil {
		fmt.Fprintln(os.Stderr, err1);
		os.Exit(1);
	}
	fmt.Println("able-mp,raklib다운로드 완료,spl다운로드중...");
	args1 := []string{"clone", "https://github.com/PocketMine/PocketMine-SPL.git", dir + "/src/spl"};
	if err2 := exec.Command(cmd, args1...).Run(); err2 != nil {
		fmt.Fprintln(os.Stderr, err2);
		os.Exit(1);
	}
}