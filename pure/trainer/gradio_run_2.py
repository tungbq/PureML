import numpy as np
import gradio as gr

notes = ["C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"]

a = {'note':None, 'octave':None, 'duration':None}

# def generate_tone(note, octave, duration):
def generate_tone(**args):
    sr = 48000
    # a4_freq, tones_from_a4 = 440, 12 * (octave - 4) + (note - 9)
    # frequency = a4_freq * 2 ** (tones_from_a4 / 12)
    # duration = int(duration)
    # audio = np.linspace(0, duration, duration * sr)
    frequency = 1000
    audio = (20000 * np.sin(audio * (2 * np.pi * frequency))).astype(np.int16)
    return (sr, audio)


# gr.Interface(
#     generate_tone,
#     [
#         gr.Dropdown(notes, type="index"),
#         gr.Slider(minimum=4, maximum=6, step=1),
#         gr.Textbox(type="number", value=1, label="Duration in seconds"),
#     ],
#     "audio",
# ).launch()



inputs = [
        gr.Dropdown(notes, type="index"),
        gr.Slider(minimum=4, maximum=6, step=1),
        gr.Textbox(type="number", value=1, label="Duration in seconds"),
    ]

print(inputs)

gr.Interface(
    generate_tone,
    inputs,
    "audio",
).launch()