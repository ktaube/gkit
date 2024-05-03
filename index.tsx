#! /usr/bin/env bun

import { Text, render } from "ink";

import Groq from "groq-sdk";

const groq = new Groq();

const chatCompletion = await groq.chat.completions.create({
  messages: [
    { role: "user", content: "Explain the importance of low latency LLMs" },
  ],
  model: "mixtral-8x7b-32768",
});

const App = () => {
  return <Text color="green">{chatCompletion.choices[0].message.content}</Text>;
};

render(<App />);
