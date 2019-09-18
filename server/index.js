let express = require('express');
let bodyParser = require('body-parser');
let guid = require('uuid/v4');

let app = express();

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

let mascotMap = {};

class Mascot {
    constructor(name, description) {
        this.name = name;
        this.description = description;
        this.id = guid();
    }
};

app.get('/:id', (req, res, next) => {
    let m = mascotMap[req.params.id];
    console.log(m)
    if (m == null) {
        return res.status(404).send();
    }

    return res.json(m);
});

app.post('/', (req, res, next) => {
    console.log('post');
    let mascot = new Mascot(req.body.name, req.body.description);
    mascotMap[mascot.id] = mascot;
    res.status(201).json(mascot);
});

app.put('/:id', (req, res, next) => {
    let m = mascotMap[req.params.id];
    if (m == null) {
        return res.status(404).send();
    }

    m.description = req.body.description;
    m.name = req.body.name;

    mascotMap[m.id] = m;

    return res.json(m);
});

app.delete('/:id', (req, res, next) => {
    console.log(`DELETE ${req.params.id}`);
    var m = mascotMap[req.params.id];

    if (m == null) {
        return res.status(404).send();
    }

    delete mascotMap[m.id];
    res.status(204).send();
});

app.get('/', (req, res, next) => {
    res.json(mascotMap);
});

app.listen(process.env.NODE_PORT || 3000, () => {
    console.log(`Server running on port ${process.env.NODE_PORT || 3000}`);
});