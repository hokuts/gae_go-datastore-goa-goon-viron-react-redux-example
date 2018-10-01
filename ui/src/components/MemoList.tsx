import * as React from 'react';
import { ListGroup, ListGroupItem } from 'reactstrap';
import { Memo } from '../states/Memo'

interface MemoListProps {
  memos: Memo[]
}

const MemoList: React.SFC<MemoListProps> = ({memos}) => {
  console.log("MemoList memos", memos)
  return (
    <ListGroup>
      { memos.map((memo) => {
        return <ListGroupItem>{memo.content}</ListGroupItem>;
      }) }
    </ListGroup>
  )
}

export default MemoList