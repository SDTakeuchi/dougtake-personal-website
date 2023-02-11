from dataclasses import dataclass

@dataclass
class Endpoint():
    url: str
    method: str

GO_HOST = 'http://localhost:8080'

# post
FIND_POSTS = Endpoint(
    GO_HOST + '/api/v1/posts/{}',
    'GET',
)

#comment
CREATE_COMMENT = Endpoint(
    GO_HOST + '/api/v1/comments',
    'POST',
)

#tag
FIND_TAGS = Endpoint(
    GO_HOST + '/api/v1/tags',
    'GET',
)
