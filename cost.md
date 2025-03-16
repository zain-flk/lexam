## Costing as Per 1M/Hour

In the second experiment Each run is taking 15 secondd. → In one hour (3600 seconds), a single “slot” (i.e. one concurrent process) can execute:  3600/15 = 240  runs/hour
  
**In concurrent environment on bigger server Each run is taking 2.4 seconds.**
  → It processed 100 tasks in 4 minutes on a 4 CPU, 8GB RAM server.

      Total time = 4 minutes = 240 seconds
      Total tasks = 100
      Average time per task:  240 sec / 100 tasks = 2.4 sec per task

---

#### **1 server with 4 CPU and 8GB RAM processes 100 tasks in 240 seconds**
  **Cost Calculation for 1 Million Runs per Hour**

    Hetzner server with 4 CPU and 8GB RAM costs 7.89 dollars per month, which is about 0.011 dollars per hour
    Cost for 663 servers per hour = 663 multiplied by 0.011 = about 7.29 dollars per hour
    Cost per run = 7.29 dollars divided by 1 million = about 0.00000729 dollars per run

#### **Cost for 16 CPU, 32GB RAM Server**
  **Cost Calculation for 1 Million Runs per Hour**

    Hetzner server with A 16 CPU, 32GB RAM server is assumed to be 4 times more powerful than a 4 CPU, 8GB RAM server.
    Total cost per hour for 1 million runs = 166 multiplied by 0.0444 = about 7.38 dollars per hour
    Cost per run for 1 million runs per hour = 7.38 dollars divided by 1 million = about 0.00000738 dollars per run

---
### **Cost and Resource Calculation Table**  

| **Setup**            | **Servers Needed** | **Cost Per Hour** | **Cost Per Run** |
|----------------------|-------------------|------------------|-----------------|
| **1 Million Runs Per Hour (4 CPU, 8GB RAM, $7.89/month)** | 663 | $7.29 | $0.00000729 |
| **1 Million Runs Per Hour (16 CPU, 32GB RAM, $32/month)** | 166 | $7.38 | $0.00000738 |


**Conclusion**

Per run cost is nearly the same across setups.

---
