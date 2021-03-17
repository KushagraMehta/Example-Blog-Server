| End-Point                 | Request Type | Function call  | Auth require |
| ------------------------- | ------------ | -------------- | ------------ |
| /login                    | POST         | Login          | no           |
| /signup                   | POST         | SiginUp        | no           |
| /users/{id}               | GET          | FindUserByID   | no           |
| /users/{id}               | PUT          | PutNewPassword | Yes          |
| /users/{id}/like          | GET          | GetLikedPost   | yes          |
| /users/{id}/like/{postID} | PATCH        | PatchLike      | yes          |

| End-Point            | Request Type | Function call | Auth require |
| -------------------- | ------------ | ------------- | ------------ |
| /posts               | POST         | Create        | yes          |
| /posts/{id}          | GET          | GetPostbyID   | no           |
| /posts/top/{limit}   | GET          | GetTopPostIDs | no           |
| /posts/draft/{id}    | GET          | GetDraft      | Yes          |
| /posts/draft/{id}    | PATCH        | PatchDrafted  | Yes          |
| /posts/{postID}/tags | GET          | GetTagsOfPost | no           |
| /posts/tag/{id}      | GET          | GetPostsOfTag | no           |

| End-Point           | Request Type | Function call | Auth require |
| ------------------- | ------------ | ------------- | ------------ |
| /tags               | POST         | Create        | yes          |
| /tags/{id}          | GET          | GetTagData    | no           |
| /tags/top/{limit}   | GET          | GetTopTags    | no           |
| /tags/{postid}/{id} | DELETE       | DeleteTags    | yes          |
| /tags/add/{id}      | POST         | AttachMe      | yes          |

| End-Point          | Request Type | Function call | Auth require |
| ------------------ | ------------ | ------------- | ------------ |
| /comments/{postID} | GET          | GetComments   | no           |
| /comments/{postID} | POST         | Post          | yes          |
| /comments/{postID} | DELETE       | Delete        | yes          |
