# Snap Control System

A simple **Go-based distributed system** where a **Master** device can manage and control multiple **Snap** devices over TCP. The system also includes a web-based GUI to monitor and control connected Snaps in real time.
![Alt Text](![Screenshot (87)](https://github.com/user-attachments/assets/6212392e-78c1-491c-a17f-a96157cac075)
)


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
│   ├── main.go        # Contains the Web GUI, TCP server, and HTML integration
│   ├── master.go      # Contains reusable static logic for TCP handling or utilities
│   ├── index.html     # Web GUI template (HTML page)
│
├── Snap/
│   └── main.go           # Client logic (connects to Master and listens for commands)
