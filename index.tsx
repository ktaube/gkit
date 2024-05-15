#! /usr/bin/env bun

import { Box, Text, render } from "ink";

import Groq from "groq-sdk";
import { $ } from "bun";

const groq = new Groq();

const gitStatus = await $`git status`.quiet().text();
const gitDiff = await $`git diff`.quiet().text();

const systemMessage = `
You are a helpful assistant that helps me developing software.
You have access to "git status" and "git diff" outputs.

Describe the changes in bulletpoints.
Look at the changes as a whole. Figure out *why* the changes are implemented.

Format:
<emoji that best describes the change> <description of change>
<emoji that best describes the change> <fun fact about one of the used methods from a known library>

Example:
📍 Adds sqlite database
💭 sqlite database is the most used database in the world
`;

const chatCompletion = await groq.chat.completions.create({
  messages: [
    { role: "system", content: systemMessage },
    { role: "user", content: `> git status\n${gitStatus}` },
    { role: "user", content: `> git diff\n${gitDiff}` },
  ],
  model: "mixtral-8x7b-32768",
});

const App = () => {
  return (
    <Box display="flex" flexDirection="column" marginRight={2}>
      <Box>
        <Text>git status</Text>
      </Box>
      <Box borderStyle="single">
        <Text color="green">{gitStatus}</Text>
      </Box>
      <Box borderStyle="single">
        <Text color="red">{gitDiff}</Text>
      </Box>
      <Box borderStyle="single">
        <Text color="blue">{chatCompletion.choices[0].message.content}</Text>
      </Box>
    </Box>
  );
};

render(<App />);
