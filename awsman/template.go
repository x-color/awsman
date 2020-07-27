package awsman

const templateText = `<html>

<body>
    <h1>AWS Accounts</h1>

    <h2>Alias List</h2>
    <ul>
        {{range .Aliases}}
        <li><a href="{{.URL}}" target="_blank" rel="noreferrer noopener">{{.Alias}}</a></li>
        <ul style="list-style: none; padding-left: 1em">
            <li>Account: {{.ID}}</li>
            <li>Role: {{.Role}}</li>
        </ul>
        {{end}}
    </ul>

    <h2>Account List</h2>
    {{range .Accounts}}
    <h3>{{.ID}}</h3>
    <ul>
        {{range .Accounts}}
        <li><a href="{{.URL}}" target="_blank" rel="noreferrer noopener">{{.Role}}</a></li>
        {{end}}
    </ul>
    {{end}}
</body>

</html>`
