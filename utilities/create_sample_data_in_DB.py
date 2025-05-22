import os
from supabase import create_client, Client
from faker import Faker

# Get Supabase URL and API key from environment variables
url: str = os.environ.get("SUPABASE_URL")
key: str = os.environ.get("SUPABASE_KEY")

# Create Supabase client
supabase: Client = create_client(url, key)

# Initialize Faker for generating fake data
fake = Faker()

def generate_sample_data(num_rows=38):
  """Generates sample data for the Users table and inserts it into Supabase.

  Args:
    num_rows: The number of rows to generate (default: 38).
  """

  for _ in range(num_rows):
    name = fake.name()
    email = fake.email()
    password = fake.password()
    bio = fake.text(max_nb_chars=200) if fake.random_int(0, 1) else None
    location = f"({fake.longitude()}, {fake.latitude()})" if fake.random_int(0, 1) else None
    profile_picture_url = fake.image_url() if fake.random_int(0, 1) else None

    data = {
        'name': name,
        'email': email,
        'password': password,
        'bio': bio,
        'location': location,
        'profile_picture_url': profile_picture_url
    }

    try:
      # Insert data into Supabase
      response = supabase.table("Users").insert(data).execute()

      if response.error:
        print(f"Error inserting data: {response.error}")
      else:
        print(f"Inserted data: {data}")

    except Exception as e:
      print(f"An error occurred: {e}")

# Generate 38 rows of sample data
generate_sample_data()