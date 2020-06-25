# How To contribute

### Find a bug? 
  Open an issue
    If you can work on it
      Fork
      Create a new branch
      Create a PR
  
## Coding conventions
1. Always run go formatter with ```go fmt```.
2. Be sure to write as much tests as you can.
3. Write comments for non-obvious algorithms and logic.
4. Keep the names for variables, constants, function, structures, etc as short as possible.
5. PR will be DECLINED if any 3rd party library is added without a very good explanation or if the same with the std lib implies a huge amount of work.
6. HTTP routing:
      - Keep 1:1 handlers and routing.
      - The name of the handler must end with the suffix "handler" i.e. MultiGetHandler(w http.ResponseWriter, r *http.Requesst) {}
      - Allow only one HTTP method.
      - No 3rd party routing libs.
      - Middlewares are allowed and welcome.
7. Avoid global variables.
8. Use internal package for shared and core logic
9. Logging is always in standard output -> use logs package, avoid using fmt package for logging.
10. Benchmark testing is welcome.

## Branching and commit messages
1. Branch name should express the intent of the change or use the same name that is in the issue.
2. Commits messages should express the change and the impact if there is a code change. Tagging with "feature", "fix", "improvement" can help, but is 
not mandatory.
3. Merge squash into master will be consider in every PR.

## Language is english along all the project.
