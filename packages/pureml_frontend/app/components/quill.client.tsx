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
        setContent(quill.root.innerHTML);
      });
    }
  }, [quill]);

  return (
    <div className="h-4/5 pr-6 text-slate-600">
      <div ref={quillRef} />
    </div>
  );
}
