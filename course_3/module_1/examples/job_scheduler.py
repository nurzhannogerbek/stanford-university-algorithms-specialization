class Job:
    """
    Represents a job with a unique ID, weight, and length.
    Attributes:
        job_id (int): The unique identifier for the job.
        weight (float): The weight or importance of the job.
        length (float): The duration required to complete the job.
    """
    def __init__(self, job_id, weight, length):
        if length <= 0:
            raise ValueError("Job length must be positive.")
        if weight < 0:
            raise ValueError("Job weight cannot be negative.")
        self.job_id = job_id
        self.weight = weight
        self.length = length

    @property
    def weight_to_length_ratio(self):
        """
        Calculates the weight-to-length ratio of the job.
        Returns:
            float: The weight-to-length ratio (weight / length).
        """
        return self.weight / self.length

    def __str__(self):
        """
        String representation of the job.
        Returns:
            str: The job details in a readable format.
        """
        return f"Job(id={self.job_id}, weight={self.weight}, length={self.length})"


class JobScheduler:
    """
    Manages the scheduling of jobs using a greedy algorithm.
    Attributes:
        jobs (list[Job]): A list of jobs to be scheduled.
    """
    def __init__(self, jobs):
        """
        Initializes the scheduler with a list of jobs.
        Args:
            jobs (list[Job]): A list of Job objects.
        """
        self.jobs = list(jobs)

    def schedule_jobs(self):
        """
        Schedules jobs based on their weight-to-length ratio in descending order.
        Returns:
            list[Job]: A list of jobs in the optimal execution order.
        """
        self.jobs.sort(key=lambda job: job.weight_to_length_ratio, reverse=True)
        return self.jobs

    def calculate_weighted_completion_time(self):
        """
        Calculates the weighted sum of completion times for the scheduled jobs.
        Returns:
            float: The weighted sum of completion times.
        """
        total_completion_time = 0
        current_time = 0

        for job in self.jobs:
            current_time += job.length  # Add job length to the current time.
            total_completion_time += job.weight * current_time  # Add weighted completion time.

        return total_completion_time


# Example usage.
if __name__ == "__main__":
    # Define a list of jobs with job_id, weight, and length.
    jobs = [
        Job(job_id=1, weight=3, length=2),
        Job(job_id=2, weight=1, length=3),
        Job(job_id=3, weight=2, length=1)
    ]

    # Initialize the job scheduler.
    scheduler = JobScheduler(jobs)

    # Schedule the jobs.
    scheduled_jobs = scheduler.schedule_jobs()
    print("Scheduled jobs (in execution order):")
    for job in scheduled_jobs:
        print(job)

    # Calculate the weighted sum of completion times.
    total_weighted_completion_time = scheduler.calculate_weighted_completion_time()
    print(f"\nWeighted sum of completion times: {total_weighted_completion_time}")
