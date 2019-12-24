# Tatsumaki Go 
[![CircleCI](https://img.shields.io/circleci/build/github/tatsuworks/tatsumaki-go/master?logo=circleci&style=flat-square&token=ef3f8f3985ba716ddabea3c9599f19ccd7307a55)](https://circleci.com/gh/tatsuworks/tatsumaki-go)
[![Code Quality](https://img.shields.io/codacy/grade/444a66fab315470a98dc427bf0e6ef4f?logo=codacy&style=flat-square)](https://www.codacy.com/manual/hassieswift621/tatsumaki-go?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=hassieswift621/tatsumaki-go&amp;utm_campaign=Badge_Grade)
[![Documentation](https://img.shields.io/badge/godoc-reference-blue?logo=go&style=flat-square)](http://godoc.org/github.com/hassieswift621/tatsumaki-go)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/hassieswift621/tatsumaki-go?logo=go&style=flat-square)](https://github.com/hassieswift621/tatsumaki-go/releases)

A Go wrapper for Tatsumaki's API, one of the most popular bots on Discord.

If you have any queries about the API or the bot, please visit <https://tatsumaki.xyz/>,
as I am not part of the Tatsumaki team.

## How do I get an API key
To get an API key, run the following command with Tatsumaki in a server: ``t!apikey``

## Tutorial
To use the lib, first create an instance of the Tatsumaki client which will be 
your interface to interact with the API.

Then, call your desired endpoint methods on the client.

```go
// Create client.
tatsumakiClient := tatsumakigo.New("YOUR API TOKEN")

// Get a user.
user, err := tatsumakiClient.User("User ID")

if err == nil {
	// Do stuff with the user response here.
	fmt.Println("User's required XP until next level up: %d", user.LevelProgress.RequiredXp)
}
...
// Get a user using a context.
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
user, err = tatsumakiClient.UserWithContext(ctx, "User ID")
...
```

## License
Copyright &copy;2019 Hassie.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

<http://www.apache.org/licenses/LICENSE-2.0>

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
