# Installs signal_loss on Windows (PowerShell): downloads the release binary,
# puts it in %LOCALAPPDATA%\signal_loss\bin and adds that folder to your user PATH.
#
#   irm https://raw.githubusercontent.com/aeon022/signal_loss/main/install.ps1 | iex
#
# $env:VERSION = "v0.4.1" picks a version first.
$ErrorActionPreference = "Stop"
$repo = "aeon022/signal_loss"
$arch = if ($env:PROCESSOR_ARCHITECTURE -eq "ARM64") { "arm64" } else { "amd64" }

$tag = $env:VERSION
if (-not $tag) {
  $tag = (Invoke-WebRequest "https://github.com/$repo/releases/latest" -MaximumRedirection 0 -SkipHttpErrorCheck).Headers.Location -replace '.*/', ''
}
$name = "signal_loss_$($tag.TrimStart('v'))_windows_$arch"
$dest = Join-Path $env:LOCALAPPDATA "signal_loss\bin"
$tmp = Join-Path ([IO.Path]::GetTempPath()) $name

New-Item -ItemType Directory -Force $dest, $tmp | Out-Null
Invoke-WebRequest "https://github.com/$repo/releases/download/$tag/$name.zip" -OutFile "$tmp.zip"
Expand-Archive "$tmp.zip" -DestinationPath $tmp -Force
Copy-Item "$tmp\signal_loss.exe" $dest -Force

$path = [Environment]::GetEnvironmentVariable("Path", "User")
if (($path -split ";") -notcontains $dest) {
  [Environment]::SetEnvironmentVariable("Path", "$path;$dest", "User")
  Write-Host "Added $dest to your PATH - open a new terminal."
}
Write-Host "Installed signal_loss $tag. Play with:  signal_loss"
