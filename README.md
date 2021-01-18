Recipe Stats Calculator
====
This code is optimized for the given dataset. It is not flexible for other JSON formats but easily changes for support. The main goal is maximum performance. So this script uses one thread (I/O operations is half of the execution time, multithread slows down execution) and just ~700 memory allocations. Most of them (~500) allocates by JSON marshal/unmarshal config and report and no one allocation for scanning. So memory consumption doesn't affect by dataset size. For the given requirements  (1M postcodes, 2K recipe names) script uses ~40Mb memory, without data size correlation. For comparison, unmarshaling by std lib the whole file and calculating takes 1,5Gb memory and 6 times more execution time, also millions of memory allocations. 

Setup
-----
Requirements `Docker` `go >= v1.12`

```
make
```

####Config file format
```
{
  "postcode_counter" : {
    "postcode": "10120",
    "from": 10,
    "to": 15
  },
  "search_by_name" : [
    "Potato",
    "Veggie",
    "Mushroom"
  ]
}
```

Run app
-----
```
make data=path_to_fixtures_file conf=path_to_config_file run
```
For example:

`make data=/home/user/data.json conf=/home/user/conf.json run` 

Task description
======
In the given assignment we suggest you to process an automatically generated JSON file with recipe data and calculated some stats.

Instructions
-----

1. Clone this repository.
2. Create a new branch called `dev`.
3. Create a pull request from your `dev` branch to the master branch.
4. Reply to the thread you're having with our HR department telling them we can start reviewing your code

Given
-----

Json fixtures file with recipe data. In the following format:
```
[
{
  "postcode": "10224",
  "recipe": "Creamy Dill Chicken",
  "delivery": "Wednesday 1AM - 7PM"
},
{
  "postcode": "10208",
  "recipe": "Speedy Steak Fajitas",
  "delivery": "Thursday 7AM - 5PM"
},
{
  "postcode": "10120",
  "recipe": "Cherry Balsamic Pork Chops",
  "delivery": "Thursday 7AM - 9PM"
},
{
  "postcode": "10186",
  "recipe": "Cherry Balsamic Pork Chops",
  "delivery": "Saturday 1AM - 8PM"
},
{
  "postcode": "10163",
  "recipe": "Hot Honey Barbecue Chicken Legs",
  "delivery": "Wednesday 7AM - 5PM"
},
{
  "postcode": "10213",
  "recipe": "Tex-Mex Tilapia",
  "delivery": "Friday 8AM - 7PM"
},
{
  "postcode": "10137",
  "recipe": "One-Pan Orzo Italiano",
  "delivery": "Wednesday 4AM - 7PM"
},
{
  "postcode": "10180",
  "recipe": "Chicken Sausage Pizzas",
  "delivery": "Saturday 6AM - 7PM"
},
{
  "postcode": "10148",
  "recipe": "Spanish One-Pan Chicken",
  "delivery": "Saturday 10AM - 4PM"
}
]
```

_Important notes_

1. Property value `"delivery"` always has the following format: "{weekday} {h}AM - {h}PM", i.e. "Monday 9AM - 5PM"
2. The number of distinct postcodes is lower than `1M`, one postcode is not longer than `10` chars.
3. The number of distinct recipe names is lower than `2K`, one recipe name is not longer than `100` chars.

Functional Requirements
------

1. Count the number of unique recipe names.
2. Count the number of occurences for each unique recipe name (alphabetically ordered by recipe name).
3. Find the postcode with most delivered recipes.
4. Count the number of deliveries to postcode `10120` that lie within the delivery time between `10AM` and `3PM`, examples _(`12AM` denotes midnight)_:
    - `NO` - `9AM - 2PM`
    - `YES` - `10AM - 2PM`
5. List the recipe names (alphabetically ordered) that contain in their name one of the following words:
    - Potato
    - Veggie
    - Mushroom

Non-functional Requirements
--------

1. The application is packaged with [Docker](https://www.docker.com/).
2. Setup scripts are provided.
3. The submission is provided as a `CLI` application.
4. The expected output is rendered to `stdout`. Make sure to render only the final `json`. If you need to print additional info or debug, pipe it to `stderr`.
5. It should be possible to (implementation is up to you):  
    a. provide a custom fixtures file as input  
    b. provide custom recipe names to search by (functional reqs. 5)  
    c. provide custom postcode and time window for search (functional reqs. 4)  

Expected output
---------------

Generate a JSON file of the following format:

```json5
{
    "unique_recipe_count": 15,
    "count_per_recipe": [
        {
            "recipe": "Mediterranean Baked Veggies",
            "count": 1
        },
        {
            "recipe": "Speedy Steak Fajitas",
            "count": 1
        },
        {
            "recipe": "Tex-Mex Tilapia",
            "count": 3
        }
    ],
    "busiest_postcode": {
        "postcode": "10120",
        "delivery_count": 1000
    },
    "count_per_postcode_and_time": {
        "postcode": "10120",
        "from": "11AM",
        "to": "3PM",
        "delivery_count": 500
    },
    "match_by_name": [
        "Mediterranean Baked Veggies", "Speedy Steak Fajitas", "Tex-Mex Tilapia"
    ]
}
```

Review Criteria
---

We expect that the assignment will not take more than 3 - 4 hours of work. In our judgement we rely on common sense
and do not expect production ready code. We are rather instrested in your problem solving skills and command of the programming language that you chose.

It worth mentioning that we will be testing your submission against different input data sets.

__General criteria from most important to less important__:

1. Functional and non-functional requirements are met.
2. Prefer application efficiency over code organisation complexity.
3. Code is readable and comprehensible. Setup instructions and run instructions are provided.
4. Tests are showcased (_no need to cover everything_).
5. Supporting notes on taken decisions and further clarifications are welcome.

