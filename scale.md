 To scale your manual server setup efficiently, We will need to **horizontal scaling (more servers)** and **vertical scaling (bigger servers)** while optimizing resource usage.

- **Reduce Build & Deploy Time:** Since CPU is busy for only 5 seconds out of 30, try parallelizing tasks better.
- **Optimize Memory Usage:** Each build uses 600MB RAM. If possible, reduce dependencies or optimize the TypeScript build process.
- **Deployment Strategy:**
  - **load balancer** to distribute jobs across multiple servers.

- **3. Vertical Scaling (Bigger Servers)**
- A **16 CPU, 32 GB RAM ($32/month) server** can theoretically handle **4x more** than a 4 CPU server.
  This reduces the number of required servers, but cost per server is higher.
