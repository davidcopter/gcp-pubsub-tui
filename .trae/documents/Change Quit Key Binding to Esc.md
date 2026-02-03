I will update the key binding to quit the application from `q` to `Esc`.

### **Implementation Steps**

1.  **Modify Key Handling (`internal/tui/model.go`)**
    *   Change the case checking for "q" to "esc".
    *   Keep "ctrl+c" as an alternative quit method.

2.  **Update Help Text (`internal/tui/model.go`)**
    *   Update the footer text to display "Press 'Esc' to quit" instead of "'q' to quit".

3.  **Update Documentation (`README.md`)**
    *   Reflect the key binding change in the "Interactive Controls" section.

### **Code Changes**
*   **`internal/tui/model.go`**: Update `Update()` function and `View()` function.
*   **`README.md`**: Update usage instructions.
