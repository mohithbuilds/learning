# Fuzzing Notes

## What is Fuzzing?
Fuzzing (or fuzz testing) is an automated software testing technique that involves providing invalid, unexpected, or random data as inputs to a computer program. The program is then monitored for exceptions, crashes, memory leaks, or incorrect behavior. The goal is to discover vulnerabilities and bugs that might not be found through traditional testing methods.

## Why Fuzz?
1.  **Discovering Edge Cases:** Fuzzing excels at finding unexpected behavior in edge cases that developers might not have considered during manual test case creation.
2.  **Security Vulnerabilities:** It's highly effective at uncovering security flaws like buffer overflows, injection flaws, denial-of-service vulnerabilities, and other crashes that could be exploited.
3.  **Robustness and Reliability:** Improves the overall robustness and reliability of software by ensuring it can handle malformed or unusual inputs gracefully.
4.  **Automated and Efficient:** Once set up, fuzzing can run continuously and automatically, making it an efficient way to find bugs without constant human intervention.
5.  **Complementary to Other Testing:** Fuzzing complements unit tests, integration tests, and manual testing by exploring different input spaces.

## When to Fuzz?
*   **During Development:** Integrate fuzzing early in the development cycle, especially for components that handle untrusted input (e.g., parsers, network protocols, file format readers).
*   **After Major Changes:** Run fuzzers after significant code changes or refactoring, as new bugs can be introduced.
*   **Continuous Integration/Continuous Deployment (CI/CD):** Incorporate fuzzing into your CI/CD pipeline to continuously test the application with new inputs.
*   **Security Audits:** As part of a security audit or penetration testing, fuzzing can be used to identify potential exploits.
*   **Handling Complex Input Formats:** Any code dealing with complex data formats, protocols, or external inputs (e.g., image decoders, archive extractors, JSON/XML parsers).

## How to Fuzz (General Steps):
1.  **Identify Target:** Choose a component or function to fuzz. This is often code that parses input, communicates over a network, or processes complex data.
2.  **Define Input Format:** Understand the expected input format for the target.
3.  **Generate Fuzzed Inputs:**
    *   **Random Fuzzing:** Generate completely random inputs (simple but less effective for complex formats).
    *   **Mutation Fuzzing:** Take valid inputs (seeds) and mutate them (flip bits, change characters, insert/delete data).
    *   **Generational Fuzzing:** Generate inputs based on a specification or grammar of the expected input format, introducing variations.
4.  **Execute Target with Fuzzed Inputs:** Run the target component with the generated fuzzed inputs.
5.  **Monitor for Anomalies:** Observe the target for:
    *   Crashes (segmentation faults, panics)
    *   Assertions failures
    *   Memory leaks
    *   Incorrect output or behavior
    *   Performance degradation
6.  **Report and Debug:** When an anomaly is found, log the fuzzed input that caused it. This input (often called a "crash reproducer") is crucial for debugging and fixing the bug.
7.  **Coverage-Guided Fuzzing:** Modern fuzzers often use code coverage feedback to intelligently guide input generation. They prioritize inputs that discover new execution paths, leading to more efficient bug discovery. (e.g., libFuzzer, AFL, Go's native fuzzing).

## Go Fuzzing Specifics (based on common Go practices):
*   **`go test -fuzz=. `**: Go 1.18+ includes native fuzzing support.
*   **Fuzz Targets:** Write `Fuzz` functions (e.g., `func FuzzReverse(f *testing.F)`) in `_test.go` files.
*   **Seed Corpus:** Add "seed" inputs using `f.Add(...)` in `Fuzz` function to guide the fuzzer. These are known valid inputs.
*   **Fuzz Corpus:** Go fuzzing maintains a corpus of interesting inputs that increase code coverage, stored in `testdata/fuzz/<FuzzTargetName>/`. These are discovered by the fuzzer.
*   **Monitoring:** The Go fuzzer automatically monitors for panics, infinite loops, and other runtime errors.
*   **Reproducers:** When a bug is found, the fuzzer generates a minimal reproducing input file in the fuzz corpus directory.