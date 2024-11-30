class Job:
    """
    Represents a job with a weight and length.
    """
    def __init__(self, weight, length):
        self.weight = weight
        self.length = length

    @property
    def weight_length_ratio(self):
        """
        Calculates the ratio (weight / length) for sorting.
        """
        return self.weight / self.length

    def __repr__(self):
        return f"Job(weight={self.weight}, length={self.length})"


class Scheduler:
    """
    Scheduler to minimize the weighted sum of completion times using the ratio-based greedy algorithm.
    """
    def __init__(self, jobs):
        """
        Initializes the scheduler with a list of jobs.
        :param jobs: List of Job objects.
        """
        self.jobs = jobs

    def schedule_jobs(self):
        """
        Schedules jobs based on the ratio (weight / length) in descending order.
        :return: A list of jobs in the scheduled order.
        """
        self.jobs.sort(
            key=lambda job: job.weight_length_ratio,
            reverse=True
        )
        return self.jobs

    def calculate_weighted_completion_time(self):
        """
        Calculates the weighted sum of completion times for the scheduled jobs.
        :return: The weighted sum of completion times as an integer.
        """
        total_time = 0
        completion_time = 0

        for job in self.jobs:
            completion_time += job.length
            total_time += job.weight * completion_time

        return total_time


# Example usage
if __name__ == "__main__":
    # The file path.
    file_path = "jobs.txt"

    # Reading jobs from the specified file.
    jobs = []
    try:
        with open(file_path, "r") as file:
            lines = file.readlines()
            num_jobs = int(lines[0].strip())
            for line in lines[1:]:
                weight, length = map(int, line.split())
                jobs.append(Job(weight, length))
    except FileNotFoundError:
        print(f"Error: The file '{file_path}' does not exist.")
        exit(1)
    except ValueError as e:
        print(f"Error: Invalid file format. {e}")
        exit(1)

    # Scheduling and calculating weighted completion time.
    scheduler = Scheduler(jobs)
    scheduler.schedule_jobs()
    total_weighted_completion_time = scheduler.calculate_weighted_completion_time()

    print(f"Total Weighted Completion Time: {total_weighted_completion_time}")
