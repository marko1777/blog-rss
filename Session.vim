let SessionLoad = 1
let s:so_save = &g:so | let s:siso_save = &g:siso | setg so=0 siso=0 | setl so=-1 siso=-1
let v:this_session=expand("<sfile>:p")
silent only
silent tabonly
cd ~/projects/learn/blog-rss
if expand('%') == '' && !&modified && line('$') <= 1 && getline(1) == ''
  let s:wipebuf = bufnr('%')
endif
let s:shortmess_save = &shortmess
if &shortmess =~ 'A'
  set shortmess=aoOA
else
  set shortmess=aoO
endif
badd +22 main.go
badd +116 cmd/main.go
badd +33 cmd/cmd.go
badd +8 internal/state/state.go
badd +19 sql/queries/users.sql
badd +3 .gitignore
badd +57 rss/rss.go
badd +12 sql/schema/002_feeds.sql
badd +9 sql/schema/001_users.sql
badd +8 sql/queries/feed.sql
badd +6 sqlc.yaml
badd +28 internal/database/feed.sql.go
badd +18 internal/database/users.sql.go
badd +12 internal/database/models.go
argglobal
%argdel
$argadd ~/projects/learn/blog-rss/
edit sql/queries/users.sql
let s:save_splitbelow = &splitbelow
let s:save_splitright = &splitright
set splitbelow splitright
wincmd _ | wincmd |
vsplit
1wincmd h
wincmd w
let &splitbelow = s:save_splitbelow
let &splitright = s:save_splitright
wincmd t
let s:save_winminheight = &winminheight
let s:save_winminwidth = &winminwidth
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe 'vert 1resize ' . ((&columns * 93 + 47) / 95)
exe 'vert 2resize ' . ((&columns * 1 + 47) / 95)
argglobal
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal nofen
silent! normal! zE
let &fdl = &fdl
let s:l = 1 - ((0 * winheight(0) + 26) / 53)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 1
normal! 0
lcd ~/projects/learn/blog-rss
wincmd w
argglobal
if bufexists(fnamemodify("~/projects/learn/blog-rss/internal/database/feed.sql.go", ":p")) | buffer ~/projects/learn/blog-rss/internal/database/feed.sql.go | else | edit ~/projects/learn/blog-rss/internal/database/feed.sql.go | endif
if &buftype ==# 'terminal'
  silent file ~/projects/learn/blog-rss/internal/database/feed.sql.go
endif
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 16 - ((12 * winheight(0) + 26) / 53)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 16
normal! 01|
lcd ~/projects/learn/blog-rss
wincmd w
exe 'vert 1resize ' . ((&columns * 93 + 47) / 95)
exe 'vert 2resize ' . ((&columns * 1 + 47) / 95)
tabnext 1
if exists('s:wipebuf') && len(win_findbuf(s:wipebuf)) == 0 && getbufvar(s:wipebuf, '&buftype') isnot# 'terminal'
  silent exe 'bwipe ' . s:wipebuf
endif
unlet! s:wipebuf
set winheight=1 winwidth=20
let &shortmess = s:shortmess_save
let &winminheight = s:save_winminheight
let &winminwidth = s:save_winminwidth
let s:sx = expand("<sfile>:p:r")."x.vim"
if filereadable(s:sx)
  exe "source " . fnameescape(s:sx)
endif
let &g:so = s:so_save | let &g:siso = s:siso_save
set hlsearch
nohlsearch
let g:this_session = v:this_session
let g:this_obsession = v:this_session
doautoall SessionLoadPost
unlet SessionLoad
" vim: set ft=vim :
