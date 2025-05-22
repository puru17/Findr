def remove_substring_from_file(filename, substring):
  """Removes a specified substring from each line in a text file.

  Args:
    filename: The name of the text file to modify.
    substring: The substring to remove from each line.
  """

  with open(filename, 'r') as f:
    lines = f.readlines()

  with open(filename, 'w') as f:
    for line in lines:
      f.write(line.replace(substring, ''))

# Example usage:
filename = 'coordinates.txt'  # Replace with your actual filename
substring = 'LatLng'  # Replace with the substring you want to remove

remove_substring_from_file(filename, substring)