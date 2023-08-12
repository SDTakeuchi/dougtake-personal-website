import { CustomRequest } from './client';


class CreateCommentRequest extends CustomRequest {
    body: string;

    constructor(body: string) {
        super();
        this.body = body;
    }
}


class DeleteCommentRequest extends CustomRequest {
    id: number;

    constructor(id: number) {
        super();
        this.id = id;
    }
}