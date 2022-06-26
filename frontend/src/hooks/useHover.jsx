// Usage
// function App() {
//   const [hoverRef, isHovered, text, setText] = useHoverTextToSpech();
//   return <div ref={hoverRef}>{isHovered ? "😁" : "☹️"}</div>;
// }

import { useEffect, useRef, useState } from 'react';

// Hook
export function useHoverTextToSpech(textInfo) {
  var synth = window.speechSynthesis;
  // var voices = synth.getVoices();
  // const [text, setText] = useState(textInfo);
  const [value, setValue] = useState(false);
  let utterThis = new SpeechSynthesisUtterance(textInfo);
  const ref = useRef(null);

  const handleMouseOver = () => {
    console.log('utterThis', utterThis);
    if (utterThis) {
      utterThis.lang = 'id';
      synth.resume();
      synth.speak(utterThis);
      utterThis.onstart = function (event) {
        console.log(
          'We have started uttering this speech: ' + event.utterance.text
        );
      };
    }
    setValue(true);
  };

  const handleMouseOut = () => {
    // stop
    if (utterThis) {
      // utterThis.onend = function (event) {
      //   console.log(
      //     'Utterance has finished being spoken after ' +
      //       event.elapsedTime +
      //       ' seconds.'
      //   );
      // };
      // pause
      synth.pause();
      utterThis.onpause = function (event) {
        console.log('Speech paused after ' + event.elapsedTime + ' seconds.');
      };
    }
    setValue(false);
  };

  useEffect(
    () => {
      const node = ref.current;
      if (node) {
        node.addEventListener('mouseover', handleMouseOver);
        node.addEventListener('mouseout', handleMouseOut);
        return () => {
          node.removeEventListener('mouseover', handleMouseOver);
          node.removeEventListener('mouseout', handleMouseOut);
        };
      }
    },
    [ref.current] // Recall only if ref changes
  );

  // useEffect(() => {
  //   if (text) {
  //     utterThis = new SpeechSynthesisUtterance(text);
  //   } else {
  //     utterThis = undefined;
  //   }
  // }, [text]);

  useEffect(() => {
    if (textInfo) {
      utterThis = new SpeechSynthesisUtterance(textInfo);
    } else {
      utterThis = undefined;
    }
  }, [textInfo]);

  return [ref, value];
}
