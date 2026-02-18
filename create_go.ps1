# 1. 自动查找当前目录下数字开头的子目录，找到最大数字
$currentDirs = Get-ChildItem -Directory -ErrorAction SilentlyContinue
$maxNumber = 0

foreach ($dir in $currentDirs) {
    if ($dir.Name -match '^\d+') {
        $num = [int]$matches[0]
        if ($num -gt $maxNumber) {
            $maxNumber = $num
        }
    }
}

$newNumber = $maxNumber + 1
$formattedNumber = $newNumber.ToString("00")

# 2. 提示用户输入目录名称
$name = Read-Host "请输入项目名称"

# 3. 检查输入是否为空
if ([string]::IsNullOrWhiteSpace($name)) {
    Write-Host "错误：名称不能为空。" -ForegroundColor Red
    exit
}

# 4. 组合成最终目录名
$dirName = "$formattedNumber$name"

# 5. 在当前目录下创建子目录
try {
    New-Item -Path ".\$dirName" -ItemType Directory -Force | Out-Null
    Write-Host "目录 '$dirName' 准备就绪..." -ForegroundColor Cyan
}
catch {
    Write-Host "创建目录失败: $_" -ForegroundColor Red
    exit
}

# 6. 拼接文件路径
$filePath = Join-Path -Path ".\$dirName" -ChildPath "main.go"

# 7. 创建 main.go 并写入 "package main"
try {
    "package main" | Set-Content -Path $filePath -Encoding UTF8
    Write-Host "成功！已创建文件: $filePath" -ForegroundColor Green
    Write-Host "内容已写入: package main" -ForegroundColor Green
}
catch {
    Write-Host "写入文件失败: $_" -ForegroundColor Red
}
