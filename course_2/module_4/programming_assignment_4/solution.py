from concurrent.futures import ProcessPoolExecutor


class TwoSumSolver:
    def __init__(self, file_path, lower_bound, upper_bound):
        """
        Initializes the TwoSumSolver class with the file path and range bounds.

        Args:
            file_path (str): Path to the input file containing integers.
            lower_bound (int): The lower bound of the target range.
            upper_bound (int): The upper bound of the target range.
        """
        self.file_path = file_path
        self.lower_bound = lower_bound
        self.upper_bound = upper_bound
        self.numbers = set()

    def load_data(self):
        """
        Loads integers from the file into a set to ensure uniqueness.

        Raises:
            Exception: If the file cannot be read or parsed.
        """
        try:
            with open(self.file_path, 'r') as file:
                self.numbers = {int(line.strip()) for line in file}
        except Exception as e:
            raise Exception(f"Error loading data from file: {e}")

    def count_valid_targets(self, numbers_chunk):
        """
        Counts valid target values for a specific chunk of numbers.

        Args:
            numbers_chunk (list): A subset of the full numbers set.

        Returns:
            set: A set of valid target values for the given chunk.
        """
        targets = set()
        for num in numbers_chunk:
            for t in range(self.lower_bound, self.upper_bound + 1):
                complement = t - num
                if complement in self.numbers and complement != num:
                    targets.add(t)
        return targets

    def count_target_values(self):
        """
        Counts the number of distinct target values t in the range [lower_bound, upper_bound]
        such that t = x + y, where x and y are distinct numbers from the input.

        Returns:
            int: The count of distinct target values.
        """
        # Divide the numbers into chunks for parallel processing.
        num_processes = 8  # Number of processes to use.
        numbers_list = list(self.numbers)
        chunk_size = len(numbers_list) // num_processes
        chunks = [
            numbers_list[i * chunk_size:(i + 1) * chunk_size]
            for i in range(num_processes)
        ]
        # Ensure the last chunk includes all remaining numbers.
        if len(numbers_list) % num_processes != 0:
            chunks[-1].extend(numbers_list[num_processes * chunk_size:])

        # Use ProcessPoolExecutor for parallel computation.
        with ProcessPoolExecutor(max_workers=num_processes) as executor:
            results = executor.map(self.count_valid_targets, chunks)

        # Combine all results into a single set of targets.
        all_targets = set()
        for result in results:
            all_targets.update(result)

        return len(all_targets)


if __name__ == "__main__":
    # File path to the input data.
    file_path = "2sum.txt"
    lower_bound = -10000
    upper_bound = 10000

    # Initialize the TwoSumSolver.
    solver = TwoSumSolver(file_path, lower_bound, upper_bound)

    # Load data from the file.
    try:
        solver.load_data()
    except Exception as e:
        print(f"Error: {e}")
        exit(1)

    # Compute the number of distinct target values.
    result = solver.count_target_values()
    print(f"The number of target values in the range [{lower_bound}, {upper_bound}] is: {result}")
