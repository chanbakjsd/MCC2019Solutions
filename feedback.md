While it's readable, this feedback is more readable when parsed as Markdown.
For easier viewing, you can read it as Markdown on GitHub at: https://github.com/lutherwenxu/MCC2019Solutions/blob/master/feedback.md

# Suggestion
1. On the "Submit All Answers" button, submit only inputs that have been filled. I once filled in test case 1, 2, 3 and hit that button to be greeted with two test case failures.
2. Use a more rational input size (character count). While I can understand the logic behind using a large input, pasting the input was terrible. I use Vim and when I hit the paste key, my Vim was just running in paste mode and it ended up wasting precious minutes for me. Palindrome is a good example in which is hard to do manually but the input size is rational.
3. Consider providing a more common in code competition format of input. By that, I mean space-separated characters that can just be passed in to stdin. I do understand that you want to make it easier for new-comers to participate but it just makes the whole experience worse for me as Go's array declaration was different as shown below:

    Go Style:
    ```
    someArray := []int{1, 2, 3, 4}
    ```
    Provided Style:
    ```
    var someArray = [1, 2, 3, 4]
    ```

    I do understand that the latter is more common and you would not be providing the former but combined with point number 2, it simply wasted my time. If I were to have the standard space-separated input, I could have just used a scanner to parse the input instead of dealing with the declaration of other languages.

    As my friend pointed out, this looks like I'm requesting for only space-separated input but in reality, I'm requesting that you provide both at once. More experienced participants don't have to deal with this and less experienced participants can copy paste like they do currently.
5. Consider open sourcing SimpleCMS. It can help people spot bugs more easily and maybe even help develop it more! You should probably also make the user ID random instead of incremented as I could easily scrap information currently (no access control for the users page).

# Notes
* The GitHub repository was created after the competition to prevent other participants from copying it. You can look at the commit time.
* I was not able to start a new repl.it instance with the GitHub repo for some reason so I included links to my files where GitHub highlighted them acceptably.

# Thanks
The competition was awesome. Thanks for reading my rant too! :)
