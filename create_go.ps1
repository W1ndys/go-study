# 1. 提示用户输入目录名称
$dirName = Read-Host "请输入项目名称(将作为文件夹名)"

# 2. 检查输入是否为空
if ([string]::IsNullOrWhiteSpace($dirName)) {
    Write-Host "错误：名称不能为空。" -ForegroundColor Red
    exit
}

# 3. 在当前目录下创建子目录
# 修改点：去掉了 $newItem = ...，并添加了 | Out-Null 以抑制默认输出
try {
    New-Item -Path ".\$dirName" -ItemType Directory -Force | Out-Null
    Write-Host "目录 '$dirName' 准备就绪..." -ForegroundColor Cyan
}
catch {
    Write-Host "创建目录失败: $_" -ForegroundColor Red
    exit
}

# 4. 拼接文件路径
$filePath = Join-Path -Path ".\$dirName" -ChildPath "main.go"

# 5. 创建 main.go 并写入 "package main"
try {
    "package main" | Set-Content -Path $filePath -Encoding UTF8
    Write-Host "成功！已创建文件: $filePath" -ForegroundColor Green
    Write-Host "内容已写入: package main" -ForegroundColor Green
}
catch {
    Write-Host "写入文件失败: $_" -ForegroundColor Red
}
