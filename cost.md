## Costing as Per 1M/Hour and 10M/Day 

**In the POC Each run is taking 30 seconds.**  
  → In one hour (3600 seconds), a single “slot” (i.e. one concurrent process) can execute:  3600/30 = 120  runs/hour

**Concurrency is limited by resources.**  
  Based on experiment:  
  - A **Hetzner server (4 CPU, 8GB RAM, $7.59)** behaves similarly to the $80 server: **16 concurrent processes (1920 runs/hour).**  
  - A **16 CPU, 32GB RAM server (cost: $32)** is assumed to scale linearly. Since a 4 CPU, 8GB server can run 16 processes, a 16 CPU, 32GB server (which is 4× larger) can run roughly **64 processes concurrently**.  
    → Runs per such server per hour: 64 × 120 = **7680 runs/hour**.

---

### **How Many Servers to Achieve 1 Million Runs per Hour?**

   **16 CPU, 32GB Server ($32, 64 processes):**  

   - This can run **64 processes** concurrently, so:  
     64 × 120 = **7,680 runs per hour per server.**  
   - To reach 1,000,000 runs/hour:  
     1,000,000 ÷ 7,680 ≈ **131 servers** (rounding up).  
   - Total cost per hour:  
     131 × $32 = **$4,192/hour.**  
   - Cost per run:  
     $32 ÷ 7,680 ≈ **$0.00417 per run.**

---

## **🔹 Summary Table (1M Runs per Hour)**  

| **Server Type**               | **Runs/Server/Hour** | **Servers Needed** | **Total Cost/Hour** | **Cost per Run**     |
|-------------------------------|----------------------|--------------------|---------------------|----------------------|
| Hetzner (4 CPU, 8GB, \$7.59)    | 1,920                | 521                | \$3,954             | \$0.00395            |
| 16 CPU, 32GB (\$32)           | 7,680                | 131                | \$4,192             | \$0.00417            |

---

## **🔹 Summary Table (10M Runs per Day)**  

| **Server Type**               | **Runs/Server/Hour** | **Servers Needed** | **Total Cost/Day** | **Cost per Run**   |
|-------------------------------|----------------------|--------------------|--------------------|-------------------|
| Hetzner (4 CPU, 8GB, $7.59)   | 1,920                | 217                | **$1,648**         | **$0.00395**      |
| 16 CPU, 32GB Server ($32)     | 7,680                | 55                 | **$1,760**         | **$0.00417**      |

---
---


### **Assumptions & Notes**
- **Each run takes 30 seconds.** Thus, a single concurrent “slot” yields 120 runs/hour.
- **Concurrency is determined by available resources.** The numbers above assume linear scaling with CPU/RAM.
- **Cost estimates** are per hour; actual costs may vary based on provider billing.
- **Overheads and inefficiencies** (e.g., networking, startup time, etc.) are not factored in and could affect real-world numbers.

---

### Not Experimented, Execution Time taken from Above POC
#### **AWS Lambda with Firecracker-based micro VMs follows a pay-per-use pricing model.**

##### **Each function execution costs based on memory and CPU usage.** 

- Memory cost: 0.0000166667 dollars per GB-second  
- CPU cost: 0.0000000021 dollars per vCPU-millisecond  

##### **For a function using 600MB (0.6GB) and running for 30 seconds:**

- Memory cost per run: 0.6 * 30 * 0.0000166667 = 0.0003 dollars  
- CPU cost per run: 30,000 milliseconds * 0.0000000021 = 0.000063 dollars  
- Total cost per run: 0.0003 + 0.000063 = 0.000363 dollars  

##### **For 10 million runs in a day:**

- Total cost: 10,000,000 * 0.000363 = 3,630 dollars per day  