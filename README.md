# zplank_week7_assign

Automated Code Generation: 
Automated code generation, such as using generate or jennifer, allow for the reduction of boilerplate code that is repeated throughout an application, resulting in a simplified code generation process and maintainability of code. I researched and read about using the generate command and can see how this could be beneficial in the Testing Go with Statistics assignment. For that assignment, I created multiple files to calculate the correlation between each pair in the Anscombe Quartet dataset, but instead could leverage the go generate command to create an overall correlation function and work in annotations to the code in order to apply the generate command to return results for each pair.

See the generate_data.go and main.go code files to see how a function and data can be defined and then automatically generate an outcome using the generate command based on the change in underlying data.

AI-assisted programming:
Github Copilot works by auto-suggesting code as you beging writing (reminds me the auto-fill function on iPhones when typing messages). It also can suggest code by simply describing what the code should be doing by using natural language. Github Copilot can be used in a variety of scenarios (across multiple languages) and I can see it being an incredibly useful tool to speed up coding times as well as keeping the code on track and aligned with the end-users goal. This tool could help with the Testing Go with Statistics assignment by either starting to code and letting Copilot suggest the next steps, or providing it with a summary of what you are looking to do and seeing the suggestions it provides. 

AI-generated code:
AI generated code can be useful when needing to obtain simple code for an assignment. I can see where this could be difficult to use since the LLM agent will not have access to the underlying data, unless a small subset that the user is able to provide, and can only offer so much. I use ChatGPT sometimes for work when I need help to determines what functions to use or how to use them in order to meet my end goal, however, it requires some adaptation of the code in order for it to fit within my needs. When working with ChatGPT for this assignment, I found it provided the code I needed but its original output was more complex than it needed to be but after a couple back and forths, ChatGPT ended on similar code that I did orginally for the correlations. However, when trying to work with it to get the linear regression, I had a bit of trouble when trying to get it to adapt the code from using 1 set of x and y variables, to finding the linear regression between the entire Anscombe Quartet dataset. I could see where this could be challenging when wanting to use this for more complex code requirements. What I did find to be the most useful when using ChatGPT was providing the chat with the code to calculate the correlation for a set of x and y variables, and ask it to create a function and test file. I could see using AI-generated code to create functions and adapt current code to be more efficient, rather than asking it to start from scratch. 


The example conversation with ChatGPT can be going in the llm_agenet_convo.txt file or through this link: https://chat.openai.com/share/d71e55be-5ea8-456a-853f-74f342946f53. 

---------------------------------------------

The llm_generate_code folder includes the below files that the ChatGPT conversation provided: 

  llm_correlations_1.go - this was the original output when asked to create the correlation between 2 sets of input data 

  llm_correlations_2.go - I then asked the chat to modify the code to provide a more simple solution and provided one set of the x and y pairs to build the correlation for. This was the result 

  llm_regression_1.go - From there, I wanted to test to see if it could produce something a bit more complex and complete the regression models for the same x and y inputs. Again, it was more complex than my original code and I did have trouble installing some of the packages it suggested. 

  llm_regression_2.go - With this second version, I wanted to see if it would suggest ways to adapt the code to use multiple x and y inputs, so I could obtain the regression for the entire Anscome Quartet dataset. It did not give me exacly what I needed so I tried one more prompt (in llm_regression_3.go below)

  llm_regression_3.go - I tried one more prompt to get the overall regression model for the dataset by providing the chat with the whole dataset and asked to create a regression model off of it. Again, it did not give me exactly what I was looking for and felt I may have been stuck in a loop where the agent was trying to key in on prior parts of the conversation when I wanted something completely new. 

  llm_correlation_function.go - After not getting what I was looking for with the regression model, I decided to see if I could provide the engine with the correlation code it developed in llm_correlations_2.go code and make a function and test file off of it. The original output provided the code back as is and suggested renaming it to be a test.go file and running the test prompt. This was not what I was looking for, so tried prompting the agent in another way and this function file was created. 

  llm_correlation_function_test.go - this was the test file that the agent suggested to go along with the final function file above. 

---------------------------------------------

Overall, I can see where each of the above AI assisted/generated code programs can help programmers be more efficient, but I do think they all have different use cases and it still takes input from programmers to be able to produce code that works. Automated Code Generation can be most useful for maintainability and repeatability of code and can help to standardize code across multiple aspects, and while it can speed up usability, I do not think using this could drastically reduce staffing of program engineers. Similarly, the use of AI-assisted programming can help speed up coding times and provide suggestions that may make code more effective and efficient, however, it will still require the use of programmers to begin working with the code and ensuring the suggestions make sense for the end goal. AI-generated code I found to be incredibly useful for when I provided it with a simple dataset and asked for something very specific, but I do not think this is scalable to use for large datasets or code that requires more logic. It took multiple prompts to get what I was looking for and I can see where a programmer with experience would likely find it easier to just code what they were trying to do instead of prompting ChatGPT multiple times to get what they are looking for. Where I do find it beneficial though, is creating supplemental code or test files off of existing code. When providing it a starting point and asking to adapt it to be something specific, it was much easier and more effective than starting from scratch with the agent. 



