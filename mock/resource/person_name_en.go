package resource

import (
	"strings"
	"time"
)

var (
	personNameEn = []string{
		"Aaron", "Abe", "Abigail", "Abraham", "Adah", "Adam", "Adams", "Adela", "Adolf", "Aileen", "Alexander", "Alice", "Allan", "Andrew", "Ann",
		"Anna", "Anne", "Antonia", "Antony", "Augustine", "Austin", "Bailey", "Baker", "Barton", "Betty", "Bill", "Blair", "Bob", "Bonnie", "Brieuse",
		"Brown", "Bruce", "Buck", "Byron", "Camp", "Carl", "Carmen", "Carter", "Catherine", "Cathleen", "Cathy", "Charles", "Charley", "Churchill",
		"Claire", "Clark", "Cole", "Constantine", "Daisy", "Dale", "David", "Dean", "Diana", "Dodge", "Donald", "Donna", "Douglas", "Edward", "Eileen",
		"Elizabeth", "Evan", "Eve", "Ford", "Fox", "Francis", "Frank", "Franklin", "Gates", "George", "Georgia", "Gill", "Green", "Ham", "Hamlet", "Hawk",
		"Helen", "Henley", "Hilary", "Hill", "Hugo", "Hunter", "Jack", "Jackson", "James", "Jane", "Jesse", "Jessica", "Jim", "John", "Johnson", "Jordan",
		"Joseph", "Julia", "Juliet", "Karl", "Katherine", "Kathleen", "Kathryn", "Kathy", "Kent", "Lane", "Laura", "Lewis", "Lincoln", "Linda", "Lucia",
		"Lucy", "Mac", "MacDonald", "Madeline", "Magdala", "Main", "Margaret", "Maria", "Marie", "Marlin", "Martha", "Martin", "Mary", "Max", "Maximilian",
		"Maxwell", "Michael", "Michel", "Mike", "Molly", "Morgan", "Moses", "Nancy", "Newton", "Noah", "Norman", "Norton", "Orlando", "Oscar", "Owen", "Paul",
		"Pete", "Peter", "Philip", "Pike", "Reynold", "Richard", "Robert", "Robin", "Robinson", "Roger", "Roland", "Rudolf", "Rudolph", "Rupert", "Ruth",
		"Sam", "Samantha", "Samuel", "Sean", "Shaw", "Shawn", "Shirley", "Sidney", "Sidon", "Simon", "Smith", "Solomon", "Sophia", "Sophie", "Stephen",
		"Steve", "Steven", "Stevenson", "Strong", "Susan", "Taylor", "Thomas", "Tobias", "Toby", "Tom", "Tommy", "Toni", "Tony", "Tyler", "Victor",
		"Victoria", "Walter", "Washington", "Wassa", "Wendy", "Wessyns", "White", "William", "Williams", "Wood", "Xavier", "York", "Zoe",
	}
	personLastNameEn = []string{
		"Aaron", "Abby", "Abel", "Abelard", "Abigail", "Abraham", "Ada", "Adalheid", "Adam", "Adelaide", "Adeline", "Adrian", "Aidan", "Ailin", "Ailsa", "Aimee",
		"Alan", "Albert", "Alex", "Alexander", "Alexandra", "Alexis", "Alfred", "Alice", "Alicia", "Alina", "Allan", "Allison", "Alva", "Alyssa", "Amanda",
		"Amber", "Amy", "Anastasia", "Andra", "Andrea", "Andrew", "Andy", "Angel", "Angela", "Angelia", "Angelina", "Angus", "Anita", "Ann", "Anna", "Anne",
		"Annie", "Anthony", "Apollo", "April", "Ariel", "Arnold", "Arthur", "Ashley", "Audrey", "August", "Austin", "Aviva", "Avivahc", "Avivi", "Barbara",
		"Barbie", "Beata", "Beatrice", "Beatrix", "Becky", "Bella", "Ben", "Benjamin", "Benson", "Bert", "Bess", "Bette", "Betty", "Bill", "Billy", "Blake",
		"Blanche", "Bob", "Bobby", "Bonnie", "Brad", "Brandon", "Brant", "Brenda", "Brendan", "Brent", "Brian", "Brianna", "Britney", "Brittany", "Brown",
		"Bruce", "Bryan", "Caleb", "Cameron", "Camille", "Candice", "Candy", "Carina", "Carl", "Carlos", "Carmen", "Carol", "Caroline", "Carrie", "Carry",
		"Cary", "Caspar", "Cassandra", "Cassie", "Catherine", "Cathy", "Cecil", "Charlene", "Charles", "Charlotte", "Chelsea", "Cheney", "Cher", "Cherry",
		"Cheryl", "Chloe", "Chris", "Christian", "Christina", "Christine", "Christopher", "Christy", "Cinderella", "Cindy", "Clair", "Claire", "Clark",
		"Claudia", "Clement", "Cliff", "Cloris", "Cody", "Cole", "Colin", "Connie", "Constance", "Cora", "Corrine", "Cosmo", "Crystal", "Cynthia", "Daisy",
		"Daniel", "Daphne", "Darcy", "Darwin", "Dave", "David", "Debbie", "Deborah", "Debra", "Demi", "Dennis", "Denny", "Derek", "Diana", "Dick", "Dolores",
		"Donald", "Donna", "Dora", "Doris", "Douglas", "Duke", "Dylan", "Eddie", "Edgar", "Edison", "Edith", "Editha", "Edmund", "Edward", "Edwin", "Eilian",
		"Elaine", "Eleanor", "Elijah", "Elizabeth", "Ella", "Ellen", "Ellie", "Elliott", "Elsa", "Elvis", "Emerald", "Emily", "Emma", "Enid", "Enterprise", "Eric",
		"Erica", "Estelle", "Esther", "Ethan", "Eudora", "Eugene", "Eva", "Evan", "Eve", "Evelyn", "Fannie", "Fanny", "Fay", "Fiona", "Flora", "Florence", "Flta",
		"Ford", "Frances", "Francis", "Frank", "Franklin", "Fred", "Frederica", "Frederick", "Frieda", "Gabriel", "Gaby", "Garfield", "Gary", "Gavin", "Geoffrey",
		"George", "Gerald", "Gillian", "Gina", "Gino", "Gladys", "Glen", "Glendon", "Gloria", "Grace", "Greta", "Gwendolyn", "Haley", "Hank", "Hannah", "Hardy",
		"Harrison", "Harry", "Hayden", "Hebe", "Heidi", "Helen", "Helena", "Hellen", "Henna", "Henry", "Hillary", "Hilton", "Howard", "Hugo", "Hunk", "Ian", "Ignace",
		"Ignativs", "Ignatz", "Ingrid", "Irene", "Iris", "Isaac", "Isabella", "Isaiah", "Ishara", "Ivan", "Ivy", "Jack", "Jackson", "Jacob", "Jacqueline", "Jade",
		"James", "Jamie", "Jane", "Janet", "Jasmine", "Jason", "Jay", "Jean", "Jeffery", "Jenna", "Jennie", "Jennifer", "Jenny", "Jeremiah", "Jerome", "Jerry", "Jesse",
		"Jessee", "Jessica", "Jessie", "Jill", "Jim", "Jimmy", "Joan", "Joanna", "Jocelyn", "Joe", "John", "Johnny", "Joliet", "Jonathan", "Jordan", "Jose", "Joseph",
		"Josephine", "Joshua", "Josie", "Joy", "Joyce", "Judith", "Judy", "Julia", "Juliana", "Julie", "June", "Justin", "Karen", "Karida", "Kate", "Katherine",
		"Katherleen", "Kathie", "Kathy", "Katie", "Katrina", "Katy", "Kay", "Kaye", "Kayla", "Keith", "Kelly", "Kelsey", "Ken", "Kennedy", "Kenneth", "Kenny", "Kerry",
		"Kevin", "Kimberly", "Kitty", "Kris", "Kristine", "Krystal", "Kyle", "Lance", "Lareina", "Larry", "Lassie", "Laura", "Lauren", "Laurent", "Lawrence", "Leander",
		"Lee", "Lena", "Leo", "Leonard", "Leopold", "Leslie", "Lillian", "Lily", "Linda", "Lisa", "Liz", "Lora", "Loren", "Lori", "Lorin", "Lorraine", "Lorry", "Louis",
		"Louisa", "Louise", "Lucia", "Lucinda", "Lucine", "Lucy", "Luke", "Lulu", "Lydia", "Lynn", "Mabel", "Madeline", "Maggie", "Mamie", "Manda", "Mandy", "Marcus",
		"Marcy", "Margaret", "Mariah", "Marilyn", "Mark", "Marks", "Mars", "Marshal", "Martha", "Martin", "Marvin", "Mary", "Mason", "Mathilda", "Matilda", "Matthew",
		"Maureen", "Mavis", "Max", "Maxine", "May", "Mayme", "Megan", "Melinda", "Melissa", "Melody", "Mercedes", "Meredith", "Mia", "Michael", "Michelle",
		"Mickey", "Mike", "Milly", "Miranda", "Miriam", "Miya", "Molly", "Monica", "Morgan", "Nan", "Nancy", "Natalie", "Natasha", "Nathan", "Nathaniel",
		"Neil", "Nelson", "Nicholas", "Nick", "Nicole", "Nikita", "Nina", "Noah", "Nora", "Norma", "Norman", "Nydia", "Octavia", "Olina", "Oliver",
		"Olivia", "Opera", "Ophelia", "Opie", "Oprah", "Oscar", "Owen", "Pamela", "Patricia", "Patrick", "Patty", "Paul", "Paula", "Pauline",
		"Pearl", "Peggy", "Peter", "Philip", "Philomena", "Phoebe", "Phyllis", "Polly", "Priscilla", "Quentin", "Quentina", "Rachel",
		"Randal", "Randall", "Randolph", "Randy", "Ray", "Raymond", "Rebecca", "Reed", "Regina", "Rex", "Richard", "Richie", "Rick", "Ricky",
		"Riley", "Rita", "Ritchie", "Robert", "Robin", "Robinson", "Rock", "Roger", "Ronald", "Rose", "Rowan", "Roxanne", "Roy", "Ruth", "Ryan",
		"Sabrina", "Sally", "Sam", "Samantha", "Sami", "Sammy", "Samson", "Samuel", "Sandra", "Sandy", "Sara", "Sarah", "Savannah", "Scarlett",
		"Scott", "Sean", "Selina", "Selma", "Serena", "Sharon", "Shawn", "Sheila", "Shelby", "Shelley", "Shelly", "Sherry", "Sheryl",
		"Shirley", "Sidney", "Sierra", "Silvia", "Simon", "Solomon", "Sonia", "Sophia", "Spark", "Spencer", "Spike", "Stacey", "Stacy",
		"Stanley", "Stella", "Stephanie", "Stephen", "Steve", "Steven", "Stewart", "Stuart", "Sue", "Sunny", "Susan", "Tamara", "Tammy", "Tanya",
		"Tasha", "Tatiana", "Ted", "Terence", "Teresa", "Terry", "Tess", "Tessa", "Thomas", "Tiffany", "Tim", "Timothy", "Tina", "Todd", "Tom",
		"Tommy", "Tony", "Tonya", "Tracy", "Tyler", "Ultraman", "Ulysses", "Ursula", "Van", "Vanessa", "Venus", "Vera", "Vern", "Vernon", "Vicky",
		"Victor", "Victoria", "Vincent", "Violet", "Virginia", "Vita", "Vivian", "Wanda", "Warner", "Warren", "Wayne", "Wendy", "Wesley", "Whitney",
		"William", "Willy", "Winnie", "Wynne", "Yolanda", "Yvette", "Yvonne", "Zachary", "Zack", "Zara", "Zelda", "Zoe", "Zoey", "Zora", "lilian",
		"lindsay", "Zatascha",
	}
)

func RandPersonNameEn() string {
	nameLength := int64(len(personNameEn))
	lastNameLength := int64(len(personLastNameEn))
	if nameLength%2 == 0 {
		nameLength -= 1
	}
	if lastNameLength%2 == 0 {
		lastNameLength -= 1
	}
	unixNano := time.Now().UnixNano()
	builder := new(strings.Builder)
	builder.WriteString(personLastNameEn[unixNano%lastNameLength] + " ")
	builder.WriteString(personNameEn[unixNano%nameLength])
	return builder.String()
}
