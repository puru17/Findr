import os
from supabase import create_client, Client

# Get Supabase URL and API key from environment variables
url: str = os.environ.get("SUPABASE_URL")
key: str = os.environ.get("SUPABASE_KEY")

# Create Supabase client
supabase: Client = create_client(url, key)

def update_table_from_file(filename, table_name, column_name):
    """
    Updates all rows in a Supabase table with values from a text file.

    Args:
      filename: The name of the text file containing the values.
      table_name: The name of the Supabase table to update.
      column_name: The name of the column in the table to update.
    """

    with open(filename, 'r') as f:
        lines = f.readlines()

    for i, line in enumerate(lines):
        # Remove any leading/trailing whitespace from the line
        value = line.strip()
        print(value)


        # Update the row in the table
        # data = supabase.table(table_name).update({column_name: value}).filter().execute()
        response = supabase.table(table_name).update({column_name: value}).eq('id', i+80).execute()
        print(i, response)
        #     print(f"Error updating row {i + 1}: {data.error}")
        # else:
        #     print(f"Updated row {i + 1} with value: {value}")

# Example usage:
filename = 'coordinates.txt'  # Replace with your actual filename
table_name = 'Users'  # Replace with your actual table name
column_name = 'location'  # Replace with the column you want to update

update_table_from_file(filename, table_name, column_name)

