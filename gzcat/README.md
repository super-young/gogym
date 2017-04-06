Use go to decompress a gzipped file.

The included file `jabberwoky.txt.gz` is gzipped. If you try to cat it, you'll get gibberish:

```
$ cat jabberwoky.txt.gz
�Xjabberwoky.txtK�
                  1�.E�Y4�l"Y}.�{�������y|�^~WL��u=B6l6�����jG��H�f�^;�5}R΁�>m��IB�f��DA(y�
�������0�}*��Yv�S,ކ�������J=1g06�k�/�!�p�,3d=.?Df���_�H)a�cȑh'U���H,�yQ�\�@=�_�P0Q�����;D� ��ؚ���{��t
t                                                                                         ���N���xx�j�#/AK��L8w"��)ˠ��]�Ұ91_�b��f���d�)W�{K�9rzQY��S1ˈ[t.s{p@L`B���
[]���  %
```

You can use the 'gzcat' command to decompress & cat the file:

```
$ gzcat jabberwoky.txt.gz
’Twas brillig, and the slithy toves
      Did gyre and gimble in the wabe:
All mimsy were the borogoves,
      And the mome raths outgrabe.

“Beware the Jabberwock, my son!
      The jaws that bite, the claws that catch!
Beware the Jubjub bird, and shun
      The frumious Bandersnatch!”

He took his vorpal sword in hand;
      Long time the manxome foe he sought—
So rested he by the Tumtum tree
      And stood awhile in thought.

And, as in uffish thought he stood,
      The Jabberwock, with eyes of flame,
Came whiffling through the tulgey wood,
      And burbled as it came!

One, two! One, two! And through and through
      The vorpal blade went snicker-snack!
He left it dead, and with its head
      He went galumphing back.

“And hast thou slain the Jabberwock?
      Come to my arms, my beamish boy!
O frabjous day! Callooh! Callay!”
      He chortled in his joy.

’Twas brillig, and the slithy toves
      Did gyre and gimble in the wabe:
All mimsy were the borogoves,
      And the mome raths outgrabe.
```

## Instructions

* [gzip.NewReader](https://golang.org/pkg/compress/gzip/#NewReader)
  * This returns the decompressed data as a reader.
  * The input io.Reader should be the compressed data.

## Example

```
$ go run zcat.go jabberwoky.txt.gz
’Twas brillig, and the slithy toves
      Did gyre and gimble in the wabe:
All mimsy were the borogoves,
      And the mome raths outgrabe.

...
```