package vo

import "slogv2/src/main/entity"

type CommentNode struct {
	Comment  entity.Comment
	Children []CommentNode
}

func NewCommentNode(comment entity.Comment) *CommentNode {
	var node CommentNode
	node.Comment = comment
	node.Children = make([]CommentNode, 0)
	return &node
}

func BuildCommentTree(comments []entity.Comment) CommentNode {
	var rootNode *CommentNode
	rootNode = new(CommentNode)
	rootNode.Comment.Cid = 0
	rootNode.Children = make([]CommentNode, 0)

	commentNodeMap := make(map[int]*CommentNode, len(comments)+1)
	commentNodeMap[0] = rootNode

	for _, comment := range comments[:] {
		commentNodeMap[comment.Cid] = NewCommentNode(comment)
	}

	for _, comment := range comments[:] {
		commentNodeMap[comment.Parent].Children = append(commentNodeMap[comment.Parent].Children,
			*commentNodeMap[comment.Cid])
	}

	return *rootNode
}
