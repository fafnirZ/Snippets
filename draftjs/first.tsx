import React, { ReactElement } from 'react'
import { Editor, EditorState, RichUtils, convertToRaw, ContentBlock, Modifier, EditorChangeType} from 'draft-js';
// import Editor from '@draft-js-plugins/editor'
// import createSideToolbarPlugin from '@draft-js-plugins/side-toolbar'
import styled from 'styled-components'

// const sideToolbarPlugin = createSideToolbarPlugin();
//const { SideToolbar } = sideToolbarPlugin;

interface Props {
  
}


const Editorwrapper = styled.div`

  color: black;

  .DraftEditor-root {
    border: 1px solid #000;
    border-radius: 5px;
    background: white;
    height: 800px;
    width: 600px;
    overflow-y: auto;
  }
  .DraftEditor-editorContainer,
  .public-DraftEditor-content {
    height: 100%;
  }

`

// blog 
// https://sendgrid.com/blog/how-we-use-draft-js-at-sendgrid/




export default function FCeditor({}: Props): ReactElement {

  const [editorState, setEditorState] = React.useState(
    () => EditorState.createEmpty(),
  );

  

  /* bold and italics key stroke inputs*/
  const handelKeyCommand = (command: any, editorState: any) => {
    const newState = RichUtils.handleKeyCommand(editorState, command);
    if (newState) {
      setEditorState(newState);
      return 'handled'
    }
    return 'not-handled'
  }

  const toggleBold = (e : React.SyntheticEvent) => {
    e.preventDefault();
    setEditorState(RichUtils.toggleInlineStyle(editorState, 'BOLD'));
  }

  /*
  getting selection state
  // https://draftjs.org/docs/api-reference-selection-state/
  // selection state and getting selections
  let selectionState = editorState.getSelection();
  let anchorKey = selectionState.getAnchorKey();
  let currentContent = editorState.getCurrentContent();
  let currentContentBlock = currentContent.getBlockForKey(anchorKey);
  */


  const handleAlignLeft = (e: React.SyntheticEvent) => {
    e.preventDefault();

    // modifying editorstate to include current block type
    // https://typescript.hotexamples.com/examples/draft-js/Modifier/setBlockType/typescript-modifier-setblocktype-method-examples.html
  


    // change to type left
    const blockTypedContent = Modifier.setBlockType(
      editorState.getCurrentContent(),
      editorState.getSelection(),
      'left'
    );
  
    const type = 'apply-block-type' as EditorChangeType;
    setEditorState(
      EditorState.push(editorState, blockTypedContent, type)
    )

  }



  /* handles changing data to json   */
  const toOBj = () => {
    // editorState : EditorState
    // getCurrentContent() => ContentState
    // getFirstBlock() => ContentBlock


    //console.log(editorState.getCurrentContent().getFirstBlock().getType())
    //console.log(editorState.getCurrentContent().convertToRaw())
    const currentContent = editorState.getCurrentContent()
    console.log(currentContent)
    const data = JSON.stringify(convertToRaw(currentContent))
    console.log(data)

    // can just store this
    // and convert raw back to contentState
  }

  // handles custom block styling for individual blocks???

  const blockStyleHandler = (contentBlock:ContentBlock): string => {
    switch (contentBlock.getType()) {
      case 'left':
        return 'align-left';
      default:
        return '';
    }
  }

  return (
    <>
      <Editorwrapper>
        <Editor
        editorState={editorState} 
        onChange={setEditorState} 
        handleKeyCommand={handelKeyCommand}
        blockStyleFn={blockStyleHandler}
        />
      </Editorwrapper>
      <div style={{display:'flex'}}>
        <button onMouseDown={toggleBold}>B</button>
        <button>I</button>
        <button onMouseDown={handleAlignLeft}>align left</button>
      </div>
      <button onClick={toOBj}>to object</button>
    </>  
  )
}
