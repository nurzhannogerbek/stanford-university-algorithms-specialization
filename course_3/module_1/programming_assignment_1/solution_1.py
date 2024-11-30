import sys


class Job:
    """
    Represents a job with a weight and length.
    """

    def __init__(self, weight, length):
        self.weight = weight
        self.length = length

    @property
    def weight_length_difference(self):
        """
        Calculates the difference (weight - length) for sorting.
        """
        return self.weight - self.length

    def __repr__(self):
        return f"Job(weight={self.weight}, length={self.length})"


class Scheduler:
    """
    Scheduler to minimize the weighted sum of completion times.
    """

    def __init__(self, jobs):
        """
        Initializes the scheduler with a list of jobs.
        :param jobs: List of Job objects.
        """
        self.jobs = jobs

    def get_sorted_jobs(self):
        """
        Returns a new list of jobs sorted based on the difference (weight - length).
        If two jobs have the same difference, prioritize by weight.
        :return: A sorted list of jobs.
        """
        return sorted(
            self.jobs,
            key=lambda job: (job.weight_length_difference, job.weight),
            reverse=True
        )

    def calculate_weighted_completion_time(self, sorted_jobs):
        """
        Calculates the weighted sum of completion times for the scheduled jobs.
        :param sorted_jobs: A list of jobs in the scheduled order.
        :return: The weighted sum of completion times as an integer.
        """
        total_time = 0
        completion_time = 0

        for job in sorted_jobs:
            completion_time += job.length
            total_time += job.weight * completion_time

        return total_time


def read_jobs_from_file(file_path):
    """
    Reads jobs from a file and creates Job objects.
    :param file_path: Path to the input file.
    :return: A list of Job objects.
    """
    jobs = []
    try:
        with open(file_path, "r") as file:
            lines = file.readlines()
            if not lines:
                raise ValueError("Input file is empty.")

            num_jobs = int(lines[0].strip())
            if num_jobs == 0:
                raise ValueError("No jobs specified in the input file.")

            for line in lines[1:]:
                try:
                    weight, length = map(int, line.split())
                    jobs.append(Job(weight, length))
                except ValueError:
                    print(f"Invalid data format in line: {line.strip()}")
        return jobs
    except FileNotFoundError:
        print(f"File not found: {file_path}")
        sys.exit(1)
    except ValueError as e:
        print(f"Error reading file: {e}")
        sys.exit(1)


if __name__ == "__main__":
    # Accept input file path as a command-line argument.
    file_path = sys.argv[1] if len(sys.argv) > 1 else "jobs.txt"

    # Read jobs from the input file.
    jobs = read_jobs_from_file(file_path)

    # Initialize the scheduler.
    scheduler = Scheduler(jobs)

    # Schedule jobs and calculate the total weighted completion time.
    sorted_jobs = scheduler.get_sorted_jobs()
    total_weighted_completion_time = scheduler.calculate_weighted_completion_time(sorted_jobs)

    # Output the result
    print(f"Total Weighted Completion Time: {total_weighted_completion_time}")
