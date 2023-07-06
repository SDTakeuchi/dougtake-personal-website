class createCommentResponse {
    
}

class createCommentRequest extends CustomRequest {
    body: string;

    constructor(body: string) {
        super();
        this.body = body;
    }
}


class deleteCommentRequest extends CustomRequest {
    id: number;

    constructor(id: number) {
        super();
        this.id = id;
    }
}