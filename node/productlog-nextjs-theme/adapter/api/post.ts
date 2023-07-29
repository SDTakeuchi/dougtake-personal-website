import { CustomRequest } from './client';

class GetPostsRequest extends CustomRequest {
    body: string;

    constructor(body: string) {
        super();
        this.body = body;
    }
}