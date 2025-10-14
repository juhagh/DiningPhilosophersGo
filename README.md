---

# 🍜 Dining Philosophers Concurrency Solution (Go)

### A concurrent solution to the classic Dining Philosophers problem using goroutines, channels, and mutex synchronization.

---

## 🧠 Overview

This project implements the **Dining Philosophers Problem** in **Go**, demonstrating concurrency control, synchronization, and deadlock prevention.

Each philosopher alternates between thinking and eating, sharing limited chopsticks (resources).
A central **host** controls access to ensure that no more than two philosophers eat at once — preventing deadlocks and starvation.

---

## ⚙️ Features

* 🧵 **Concurrent execution** using goroutines
* 🔄 **Safe synchronization** with `sync.Mutex` and `sync.WaitGroup`
* 📡 **Channel-based communication** with a request–reply system
* 🚫 **Deadlock avoidance** via centralized permission management (the host)
* 🍽️ Each philosopher eats exactly three times

---

## 🧩 How It Works

1. Each **philosopher** requests permission to eat by sending a `Request` to the host.
2. The **host** ensures that no more than two philosophers eat concurrently.
3. Once granted, the philosopher locks their two chopsticks, eats, then releases them.
4. The philosopher notifies the host that they are done, freeing a slot for others.
5. This continues until every philosopher has eaten three times.

---

## 🏃‍♂️ Running the Program

### 1️⃣ Clone this repository:

```bash
git clone https://github.com/juhagh/DiningPhilosophersGo.git
cd DiningPhilosophersGo
```

### 2️⃣ Run the program:

```bash
go run PhilosophersDinnerParty.go
```

You should see output similar to:

```
starting to eat 1
finishing eating 1
starting to eat 3
finishing eating 3
...
```

---

## 🧩 Code Structure

| File                      | Description                                                                               |
| ------------------------- | ----------------------------------------------------------------------------------------- |
| `PhilosophersDinnerParty.go`                 | Entry point — sets up philosophers, chopsticks, and the host                              |

---

## 📚 Key Concepts Demonstrated

* **Goroutines:** lightweight concurrent execution
* **Mutexes:** prevent race conditions on shared resources
* **Channels:** enable safe communication between goroutines
* **WaitGroups:** synchronize program termination
* **Deadlock prevention:** by limiting concurrent access via a host goroutine

---

## 🧱 Future Improvements

* [ ] Parameterize philosopher count and eating rounds
* [ ] Add configurable maximum concurrent eaters
* [ ] Introduce `time.Sleep()` delays for realism
* [ ] Add simple `testing` suite to verify correct behavior
* [ ] Add metrics/logging for analysis (who waits longest, etc.)
* [ ] Split into multiple files for cleaner structure

---

## 🧩 Version History

| Version  | Description                                            | Status        |
| -------- | ------------------------------------------------------ | ------------- |
| **v1.0** | Initial working solution (host-controlled concurrency) | ✅ Released    |
| **v2.0** | Planned: modular structure, configurable constants     | ⏳ In progress |

---

## 🧠 What I Learned

* Designing concurrent systems in Go
* Managing shared resources safely with `sync.Mutex`
* Avoiding deadlocks using a centralized controller
* Clean struct-based concurrency design with goroutines and channels

---

## 👨‍💻 Author

**Juha**
Backend Developer | Go, C#, Python | Building scalable systems with Docker, Kubernetes and databases

---
