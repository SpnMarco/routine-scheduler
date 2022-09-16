# The Routine Scheduler

The objective is to design a Go binary which can execute several scheduled routines in parallel. Each routine can be scheduled to a different time slot (e.g. Database cleanup every 5min, logs cleanup every 1 hour, ...). No user input is required, but the scheduler should be organized in a way that is easy for another developer to extend with other routines or change every routine time slot.

In the pkg directory you will find some mocked actions that the routine scheduler should handle:

- Database Cleanup (pkg/database/clean.go).
- Logs Cleanup (pkg/logs/clean.go).

The project can be structured as desired, the most important point being the ability to easily extend the scheduler (ex by another developer) by adding new tasks or changing the time intervals of tasks already present.

## Bonus

While the scheduler is running it should mantain a json file (it can be in the same directory)
which should be structured like this:

```json
{
  "database_clean_routine": {
    "run_no": 13,
    "run_errors": 0
  },
  "logs_clean_routine": {
    "run_no": 4,
    "run_errors": 1
  }
  ...
}
```

runs_no should be updated each time a routine finish his job.  
run_errors should be updated each time a routine returns with an error
