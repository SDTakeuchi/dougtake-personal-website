enum HttpMethods {
    Get = 'GET',
    Post = 'POST',
    Delete = 'DELETE',
    Update = 'UPDATE'
}

class Endpoint {
    name: string;
    method: HttpMethods;
    url: string;

    constructor(name: string, method: HttpMethods, url: string) {
        this.name = name;
        this.method = method;
        this.url = url;
    }
}

const apiRoot: string = 'http://localhost/api/';

const endpoints: Array<Endpoint> = [
    new Endpoint(
        'createComment',
        HttpMethods.Post,
        'v1/comment'
    ),
    new Endpoint(
        'deleteComment',
        HttpMethods.Delete,
        'v1/comment'
    )
]