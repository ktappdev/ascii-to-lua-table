<h1>ASCII to Lua Table Converter</h1>

<p>This is a simple command-line tool written in Go to convert text input into Lua table format. The tool reads input from stdin, processes it, and outputs the corresponding Lua table.</p>

<p>I will make a simple frontend for this sometime</p>

<h2>Installation</h2>

<p>To use this tool, make sure you have Go installed on your system. Then, you can clone this repository and build the executable:</p>

<pre><code>git clone https://github.com/yourusername/text-to-lua-table.git
cd text-to-lua-table
go build
</code></pre>

<h2>How It Works</h2>

<ol>
<li><strong>Read Input</strong>: The tool reads input from stdin until an empty line is encountered.</li>
<li><strong>Find Longest Line</strong>: It determines the longest line in the input to calculate padding for shorter lines.</li>
<li><strong>Convert to Lua Table</strong>: Each line is converted into a Lua table string, adding necessary padding to match the length of the longest line.</li>
<li><strong>Output</strong>: Finally, the Lua table is printed to stdout.</li>
</ol>

<h2>Contributing</h2>

<p>Contributions are welcome! If you encounter any issues or have suggestions for improvements, please feel free to open an issue or submit a pull request.</p>

