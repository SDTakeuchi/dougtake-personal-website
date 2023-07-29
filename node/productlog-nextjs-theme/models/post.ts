class Post {
    private _id: number;
    private _title: string;
    private _body: string;
    private _tagIDs: number[];
    // comments: Comment[];

    constructor(id: number, title: string, body: string, tagIDs: number[]) {
        this._id = id;
        this._title = title;
        this._body = body;
        this._tagIDs = tagIDs;
    }

    get id(): number { return this._id; }
    get title(): string { return this._title; }
    get body(): string { return this._body; }
    get tagIDs(): number[] { return this._tagIDs; }
}