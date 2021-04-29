const MarkdownIt = require('markdown-it');
const sanitizer = require('markdown-it-sanitizer');
const emoji = require('markdown-it-emoji');
const imsize = require('markdown-it-imsize');
import markdownItAnchor from 'markdown-it-anchor'
const markdownItTableOfContents = require('markdown-it-table-of-contents')
import hljs from 'highlight.js'

export interface ITocItem {
  name: string
  link: string
}

export function markdown(text: string): [string, Array<ITocItem>] {
  let anchorId = 1;
  let tocId = 1;
  let tocs: Array<ITocItem> = new Array();
  const md = new MarkdownIt({
    html:true,
    breaks:true,
    langPrefix: 'hljs ',
    highlight: function(code, lang){
      return hljs.highlightAuto(code, [lang]).value
    }
  }).use(sanitizer)
    .use(emoji)
    .use(imsize)
    .use(markdownItAnchor, {
      slugify: function(s){
        tocs.push({name: s, link: `#${anchorId.toString()}`})
        return (anchorId++);
      }
    })
    .use(markdownItTableOfContents, {
      includeLevel: [1, 2, 3],
      slugify: function(s){
        return (tocId++)
      }
    })
  return[md.render(text), tocs];
}
