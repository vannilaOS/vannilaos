# Vanilla OS - מנהל חבילות Apx

Apx הוא מנהל החבילות ברירת המחדל של Vanilla OS. הוא עטיפה סביב מנהלי חבילות מרובים להתקנת חבילות והפעלת פקודות בתוך מיכל מנוהל.

## סקירה כללית

**Apx** (/à·peks/) מספק ממשק אחיד לניהול חבילות על פני הפצות Linux שונות. הוא משתמש ב-[distrobox](https://github.com/89luca89/distrobox) כדי ליצור סביבות בטחונות שבהן ניתן להתקין חבילות ללא השפעה על מערכת המארח שלך.

## תכונות

- **תמיכה בהפצות מרובות**: התקנת חבילות מהפצות Linux שונות
- **מבוסס על קונטיינר**: חבילות מותקנות בקונטיינרים מבודדים
- **הפשטת מנהל החבילות**: ממשק אחיד עבור מנהלי חבילות שונים (apt, dnf, pacman, וכו')
- **שילוב שולחן עבודה**: ייצוא אוטומטי של רשומות שולחן עבודה לאפליקציות מותקנות
- **ניהול תת-מערכות**: יצירה וניהול מרובות תת-מערכות עם מחסניות שונות

## הערימה

- **שפה**: Go 1.25.0
- **Framework**: Vanilla OS SDK + CLI builder
- **ממשק משתמש**: Charmbracelet (bubbletea, lipgloss) לממשק טרמינל
- **קונטיינר**: שילוב distrobox

## מבנה הפרויקט

```
cmd/
  main.go              # נקודת כניסה של היישום
  main_check.go        # וריאנט אימות מחרוזות
  locales/             # תמיכה בשפות מרובות (33+ שפות)

internal/cli/
  structs.go           # הגדרות פקודות CLI
  runtime.go           # מטפלי ביצוע תת-מערכות
  subsystems.go        # ניהול מחזור חיי תת-מערכות
  stacks.go            # פעולות ערימה
  pkgmanagers.go       # פעולות מנהל חבילות

core/                  # לוגיקה עסקית ליבית (תת-מערכות, ערימות, ניהול pkg)
config/                # קבצי תצורה
distrobox/             # סקריפטי שילוב distrobox
```

## בנייה

### דרישות מוקדמות

- Go 1.25.0 או גבוה יותר
- make
- podman או docker
- git

### שלבי בנייה

```bash
# שיבוט וכניסה לתיקייה
git clone --recursive https://github.com/yourusername/Vanilla-.git
cd Vanilla-

# בנייה
make build

# התקנה לרמת המערכת
sudo make install
sudo make install-manpages

# התקנה למיקום מותאם אישית
make install PREFIX=$HOME/.local
make install-manpages PREFIX=$HOME/.local
```

## שימוש

### פקודות בסיסיות

```bash
# רישום תת-מערכות זמינות
apx subsystems list

# יצירת תת-מערכת חדשה
apx subsystems new --name=ubuntu-latest

# כניסה לתת-מערכת
apx ubuntu-latest enter

# התקנת חבילות בתת-מערכת
apx ubuntu-latest install package-name

# הפעלת פקודה בתת-מערכת
apx ubuntu-latest run command

# רישום ערימות זמינות
apx stacks list

# ניהול מנהלי חבילות
apx pkgmanagers list
```

## תלויויות

### תלויויות ישירות
- `github.com/google/uuid` - יצירת UUID
- `github.com/vanilla-os/sdk` - Framework יישום
- `gopkg.in/yaml.v2` - תצורת YAML

### תלויויות ממשק משתמש
- `github.com/charmbracelet/bubbletea` - Framework ממשק טרמינל
- `github.com/charmbracelet/lipgloss` - עיצוב טרמינל

## תיעוד

לתיעוד מפורט, בקר ב: https://docs.vanillaos.org/docs/en/apx

## תרגומים

תרום תרגומים דרך [Weblate](https://hosted.weblate.org/projects/vanilla-os/apx)

## רישיון

GNU General Public License v3.0

## מחברים

- Mirko Brombin <brombin94@gmail.com>
- Pietro di Caprio <pietro@fabricators.ltd>
- תורמי Vanilla OS

## תרומה

תרומות מתקבלות בברכה! אנא ודא:

1. תלויויות מעודכנות עם `go get`, `go mod tidy`, ו-`go mod vendor`
2. הקוד עוקב אחר קונבנציות Go
3. תרגומים מעודכנים ב-Weblate
