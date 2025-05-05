# Snap Control System

A simple **Go-based distributed system** where a **Master** device can manage and control multiple **Snap** devices over TCP. The system also includes a web-based GUI to monitor and control connected Snaps in real time.

## 🚀 Features

- Real-time TCP connection between Master and Snap devices.
- Web GUI to list all connected Snap devices.
- Ability to send a `shutdown` command to any Snap device.
- Auto-generated Snap IDs.
- Simple, clean HTML/CSS UI.

## 🛠️ Technologies Used

- Language: Go (Golang)
- Network: TCP Sockets
- Web GUI: HTML, CSS, Go `net/http` + `html/template`
- OS-level commands (shutdown) via Go `exec`

---

## 📁 Project Structure

```bash
├── Master/
│   ├── main.go           # Main logic for the Master (TCP server + Web server)
│   ├── index.html        # Web GUI template
│   └── images/           # Folder to store uploaded images (optional)
│
├── Snap/
│   └── main.go           # Client logic (connects to Master and listens for commands)
