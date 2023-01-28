import type { Dispatch, SetStateAction } from "react";
import { useEffect } from "react";
import { useQuill } from "react-quilljs";

type PropsType = {
  setContent: Dispatch<SetStateAction<string>>;
  defaultValue?: string;
};

export default function Quill({ defaultValue, setContent }: PropsType) {
  const { quill, quillRef } = useQuill();
  useEffect(() => {
    if (quill && defaultValue) {
      quill.clipboard.dangerouslyPasteHTML(defaultValue);
    }
  }, [quill]);
  useEffect(() => {
    if (quill) {
      quill.on("text-change", (delta, oldDelta, source) => {
        // console.log('Text change!');
        // console.log(quill.getText()); // Get text only
        // console.log(quill.getContents()); // Get delta contents
        // console.log(quill.root.innerHTML); // Get innerHTML using quill
        // console.log(quillRef.current.firstChild.innerHTML); // Get innerHTML using quillRef
        setContent(quill.root.innerHTML);
      });
    }
  }, [quill]);

  return (
    <div className="w-2/3" style={{ height: 300 }}>
      <div ref={quillRef} />
    </div>
  );
}
