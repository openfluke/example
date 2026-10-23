# qnn_stress_test — 30s Hexagon HTP flex

Prepare one Dense graph (`32×512×512`), then hammer `graphExecute` for **30 seconds** with a live countdown bar.

## Run

**cmd.exe** (easiest):

```bat
cd c:\git\example\welvet\qnn_stress_test
run.cmd
```

Or set PATH yourself then `go run .`:

```bat
set PATH=c:\git\welvet\native\qnn\windows_arm64;%PATH%
cd c:\git\example\welvet\qnn_stress_test
go run .
```

**PowerShell:**

```powershell
$env:PATH = "c:\git\welvet\native\qnn\windows_arm64;$env:PATH"
cd c:\git\example\welvet\qnn_stress_test
go run .
```

`0xc0000135` = DLL not on PATH (use `run.cmd` or the `set PATH=...` line above).

Ctrl+C aborts early. End screen prints ops, throughput, and a checksum so you know it really ran.
