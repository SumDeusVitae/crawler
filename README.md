# **Go Web Crawler**

A simple and efficient web crawler implemented in Go. The crawler visits pages on a specified website, extracts links, and reports visited pages while respecting concurrency limits and a maximum number of pages to crawl.

---

## **Features**

- Crawls a website starting from a base URL.
- Extracts and normalizes links from HTML pages.
- Limits concurrency using a configurable number of workers.
- Stops crawling after reaching a maximum number of pages.
- Handles invalid or non-HTML pages gracefully.
- Thread-safe operations using mutexes.

---

## **Prerequisites**

- **Go (1.19 or higher)**: Install Go from [here](https://golang.org/dl/).

---

## **Usage**

### **Run the Crawler**

```bash
go run . <baseURL> <maxConcurrency> <maxPages>
```

### **Arguments**

- `baseURL`: The starting URL for the crawl (e.g., `https://news.ycombinator.com/`).
- `maxConcurrency`: (Optional) The maximum number of concurrent requests (default: 5).
- `maxPages`: (Optional) The maximum number of pages to crawl (default: 10).

### **Example**

```bash
go run . https://news.ycombinator.com/ 5 25
```

This example starts crawling `https://news.ycombinator.com/` with a maximum of 5 concurrent workers and stops after visiting 25 pages.

---

## **Output**

The crawler logs:
- Crawled URLs.
- Errors encountered (e.g., parsing issues or HTTP errors).
- A final report showing all visited pages and their visit counts.

Sample output:

```
Starting crawl of: https://news.ycombinator.com/
crawling https://news.ycombinator.com/
crawling https://news.ycombinator.com/from?site=humanloop.com
crawling https://news.ycombinator.com/vote?id=42325011&how=up&goto=news
Error - getHTML: error status code: 503
crawling https://news.ycombinator.com/active
Maximum pages limit of 25 is reached
=====================================
REPORT for https://news.ycombinator.com/
=====================================
Found 196 internal links to news.ycombinator.com/item
Found 90 internal links to news.ycombinator.com/from
Found 90 internal links to news.ycombinator.com/user
Found 89 internal links to news.ycombinator.com/vote
...
```

---

## **Project Structure**

```
.
├── main.go                 # Entry point of the application
├── config.go               # Configuration struct and helper methods
├── crawl.go                # Core crawling logic
├── html.go                 # Get html body logic
├── normalize_url.go        # Normalize url logic
├── report.go               # Report function logic
├── README.md               # Project documentation
```

---

## **Key Functions**

- **`crawlPage`**: Recursively crawls pages, adhering to concurrency and page limits.
- **`getHTML`**: Fetches and validates the HTML content of a given URL.
- **`getURLsFromHTML`**: Extracts valid URLs from HTML pages.
- **`normalizeURL`**: Standardizes URL format for consistency.
- **`printReport`**: Prints report of crawled pages.

---

## **Installation**

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/go-web-crawler.git
   cd crawler
   ```

2. Build the project:
   ```bash
   go build
   ```

3. Run the crawler:
   ```bash
   ./crawler <baseURL> <maxConcurrency> <maxPages>
   ```

---

## **Customization**

You can modify the following settings in `configure()`:

- **Timeout**: Adjust the HTTP client's timeout duration.
- **Concurrency**: Modify the size of the `concurrencyControl` channel.

---

## **Contributing**

Contributions are welcome! Feel free to submit issues, fork the repository, and create pull requests.

---

## **License**

This project is licensed under the [MIT License](LICENSE).

---

## **Contact**

For questions or feedback, reach out via [My website](https://kskyi.netlify.app/).

---

Feel free to adjust sections like **Contact**, **License**, and **Contributing** to match your preferences! Let me know if you’d like additional content or tweaks.
