# qnn_stress_test — 30s Hexagon HTP flex

```bat
cd c:\git\example\welvet\qnn_stress_test
go run .
```

That’s it. The outer `main` finds `libwelvet_qnn.dll`, puts it on `PATH`, sets `WELVET_QNN_INPROC=1` (skips the missing `welvet_qnn_probe.exe` gate), then runs `./stress`.

Ctrl+C aborts early. End screen prints ops, throughput, and a checksum.
