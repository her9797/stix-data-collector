package main

import (
    "bytes"
    "fmt"
    "os"
    "os/exec"
    "strings"
    "time"
)

func main() {
    ticker := time.NewTicker(1 * time.Minute) // 10분마다 반복
    defer ticker.Stop()

    fmt.Println("[INFO] TAXII 엔진 시작됨. 10분마다 리포지토리 갱신을 시도합니다.")

    for {
        checkAndUpdate()

        fmt.Println("[INFO] 다음 체크까지 대기 중...\n")
        <-ticker.C
    }
}

func checkAndUpdate() {
    repoPath := "cti"
    repoURL := "https://github.com/mitre/cti.git"

    if _, err := os.Stat(repoPath); os.IsNotExist(err) {
        fmt.Println("[INFO] 리포지토리가 없음 → git clone 시작")
        cmd := exec.Command("git", "clone", repoURL, repoPath)
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        if err := cmd.Run(); err != nil {
            fmt.Println("[ERROR] git clone 실패:", err)
            return
        }
        printLastCommit(repoPath)
        return
    }

    fmt.Println("[INFO] 기존 리포지토리 존재 → 버전 체크 중...")

    localHash := gitOutput("git", "-C", repoPath, "rev-parse", "HEAD")
    remoteHash := gitOutput("git", "-C", repoPath, "ls-remote", "origin", "HEAD")
    remoteHash = strings.Split(remoteHash, "\t")[0]

    if localHash == remoteHash {
        fmt.Println("[INFO] 현재 로컬은 최신 상태입니다. pull 생략.")
    } else {
        fmt.Println("[INFO] 새 커밋 감지됨 → git pull 실행")
        cmd := exec.Command("git", "-C", repoPath, "pull")
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        if err := cmd.Run(); err != nil {
            fmt.Println("[ERROR] git pull 실패:", err)
            return
        }
        printLastCommit(repoPath)
    }
}

func gitOutput(name string, args ...string) string {
    var out bytes.Buffer
    cmd := exec.Command(name, args...)
    cmd.Stdout = &out
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        return ""
    }
    return strings.TrimSpace(out.String())
}

func printLastCommit(path string) {
    msg := gitOutput("git", "-C", path, "log", "-1", "--pretty=format:%h %s (%cd)", "--date=short")
    fmt.Printf("[INFO] 최신 커밋 정보: %s\n", msg)
}
